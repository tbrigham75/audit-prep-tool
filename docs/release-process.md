# Release process

1. Inspect `git status` and review the proposed diff.
2. Run `node tests/regression-checks.mjs` and relevant manual regression flows.
3. Confirm no credentials, local environment files, or generated artifacts are staged.
4. Update the application version, `CHANGELOG.md`, and README/documentation where applicable.
5. Build each launcher from the repository root:

   ```powershell
   $env:GOOS='windows'; $env:GOARCH='amd64'; go build -trimpath -ldflags='-s -w' -o audit-prep-tool-windows-amd64.exe launcher.go
   $env:GOOS='linux'; $env:GOARCH='amd64'; go build -trimpath -ldflags='-s -w' -o audit-prep-tool-linux-amd64 launcher.go
   $env:GOOS='darwin'; $env:GOARCH='amd64'; go build -trimpath -ldflags='-s -w' -o audit-prep-tool-macos-amd64 launcher.go
   $env:GOOS='darwin'; $env:GOARCH='arm64'; go build -trimpath -ldflags='-s -w' -o audit-prep-tool-macos-arm64 launcher.go
   ```

6. Verify the packaged Windows or Linux launcher serves `GET /__health` as `audit-prep-tool`, and verify the embedded page reports the intended `APP_VERSION`.
7. Generate and verify SHA-256 checksums; package and validate the release ZIP.
8. Commit the source and documentation, then create an annotated `vX.Y.Z` tag.
9. Publish binaries, checksums, and release notes as GitHub Release assets rather than committing binaries to the source branch.

The source version, embedded launcher payload, documentation, checksum manifest, tag, and release notes must agree before publication.
