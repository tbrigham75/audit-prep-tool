# Testing

Run the dependency-free regression checks with:

```powershell
node tests/regression-checks.mjs
go build -o build/audit-prep-tool.exe launcher.go
```

The scripted checks guard key implementation contracts in the shipped single-file UI and launcher: current version, IndexedDB persistence, evidence hashing and size validation, duplicate detection, encryption primitives, PBC synchronization for annotated evidence, deletion cleanup, explicit AI vision gating, credential-excluded backup behavior, loopback-only binding, and health endpoint support.

Manual regression testing should cover navigation and themes; workspace switching; artifact upload/edit/filtering; control/PBC/finding mappings and deletion cleanup; notes; dashboard metrics; search; activity export; annotation/redaction; JSON/CSV/control imports; report, CSV, and ZIP exports; encryption enable/change/disable; Ollama and OpenAI-compatible connector tests; and screenshot vision opt-in.

For a release, also build each target launcher, confirm its embedded application version and `GET /__health` response, calculate checksums, and validate the generated release ZIP. The v3.0.2 user-focused regression result is preserved outside Git as a release artifact and summarized in the changelog.
