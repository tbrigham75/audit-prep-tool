package main

import (
    "embed"
    "fmt"
    "io"
    "log"
    "net"
    "net/http"
    "os/exec"
    "runtime"
    "time"
)

//go:embed index.html
var assets embed.FS

const addr = "127.0.0.1:51327"
const url = "http://127.0.0.1:51327/"

func openBrowser(u string) error {
    var cmd *exec.Cmd
    switch runtime.GOOS {
    case "windows":
        cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", u)
    case "darwin":
        cmd = exec.Command("open", u)
    default:
        cmd = exec.Command("xdg-open", u)
    }
    return cmd.Start()
}

func existingInstance() bool {
    c := http.Client{Timeout: 700 * time.Millisecond}
    r, err := c.Get(url + "__health")
    if err != nil {
        return false
    }
    defer r.Body.Close()
    b, _ := io.ReadAll(io.LimitReader(r.Body, 64))
    return r.StatusCode == http.StatusOK && string(b) == "audit-prep-tool"
}

func main() {
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        if existingInstance() {
            _ = openBrowser(url)
            return
        }
        log.Fatalf("Audit Prep Tool could not start on %s: %v", addr, err)
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/__health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        _, _ = io.WriteString(w, "audit-prep-tool")
    })
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" && r.URL.Path != "/index.html" {
            http.NotFound(w, r)
            return
        }
        b, err := assets.ReadFile("index.html")
        if err != nil {
            http.Error(w, "Embedded UI unavailable", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Header().Set("Cache-Control", "no-store")
        _, _ = w.Write(b)
    })

    srv := &http.Server{Handler: mux, ReadHeaderTimeout: 5 * time.Second}
    fmt.Println("Audit Prep Tool is running locally at", url)
    fmt.Println("Keep this window open while using the app. Close it when finished.")
    _ = openBrowser(url)
    if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}
