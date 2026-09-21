# Audit Prep Tool

**Version:** 3.0.2  
**Deployment model:** Self-contained local desktop web application  
**Default launcher URL:** `http://127.0.0.1:51327/`

Audit Prep Tool is a local-first evidence workspace for audit, assessment, and compliance teams. It turns the scattered work of responding to an auditor—screenshots, request lists, control mappings, reviewer notes, findings, and exports—into a structured, reviewable audit package that stays on the user's machine by default.

Use it when you need to answer a PBC (Provided By Client) list, prepare for SOC 2, CMMC, CIS, PCI DSS, ISO 27001, NIST, or internal security reviews, or simply organize technical evidence before an assessor asks for it. Each audit engagement has its own workspace, so evidence and review context do not mingle across clients, systems, or assessment periods.

### What it helps you do

- Collect and hash screenshot evidence, document what it proves, and track review status.
- Map evidence to framework-neutral controls and auditor/PBC requests.
- Track open findings, owners, remediation, and due dates alongside the supporting evidence.
- Redact or annotate a screenshot while preserving the original and its lineage.
- Export a manifest, printable report, backup, or auditor-facing ZIP package.
- Optionally ask Ollama or an OpenAI-compatible service about the audit workspace, with screenshot sharing disabled unless explicitly enabled.

The application is intentionally packaged without Docker, Node.js, Python, npm, a database server, or a separate web server requirement for end users. The supplied platform launchers embed the application and serve it only on loopback.

## Quick start

Build or obtain the launcher for your operating system:

- Windows 64-bit: `audit-prep-tool-windows-amd64.exe`
- Linux 64-bit: `audit-prep-tool-linux-amd64`
- macOS Intel: `audit-prep-tool-macos-amd64`
- macOS Apple Silicon: `audit-prep-tool-macos-arm64`

The source repository intentionally excludes compiled launchers. Build them using the commands in [Rebuilding launchers](#rebuilding-launchers), or obtain them from a verified release package when one is available.

Keep the launcher process open while using the app. The launcher listens only on `127.0.0.1:51327` and opens the default browser.

`index.html` is also completely self-contained and can be opened directly in a modern browser, although some enterprise browser policies block `file://` applications. The packaged launcher avoids that dependency.

> The included binaries are not code-signed. Windows SmartScreen, macOS Gatekeeper, or enterprise application-control policies may warn or block unsigned binaries. For managed deployment, sign the appropriate binary with your organization's normal code-signing process.

## Core workflow

The product is organized around this audit-preparation flow:

**Audit Workspace -> Auditor/PBC Request -> Control -> Evidence -> Review -> Finding/Remediation -> Export**

An organization can maintain multiple audit workspaces without mixing evidence between engagements.

## Implementation plan and status

The feature roadmap below is the implementation plan for the major product recommendations. All items listed here are implemented in the current 3.0.x application unless explicitly marked otherwise. Future changes should update both this table and the Changelog.

| Recommendation | Implementation approach | Status |
| --- | --- | --- |
| Audit / engagement workspaces | Isolate evidence, controls, requests, findings, notes, chat, and activity under a workspace ID; add audit metadata and workspace switching. | Implemented |
| Evidence IDs and integrity hashes | Assign sequential `ART-####` IDs and calculate SHA-256 on upload and annotation output. | Implemented |
| Control-to-evidence mapping | Store control IDs on evidence and expose evidence counts/status from Controls. | Implemented |
| Tags, filters, and search | Add artifact search/filtering plus global workspace search across evidence, controls, PBC requests, and findings. | Implemented |
| Evidence status workflow | Support Collected, Reviewed, Accepted, Needs Update, and Rejected, surfaced on Home and Artifacts. | Implemented |
| Screenshot annotation/redaction | Provide highlight, arrow, blackout redaction, and undo; preserve originals by creating a new evidence version. | Implemented |
| Structured evidence notes | Add description, internal note, auditor comment, owner, system/application, collection method, tags, and mappings. | Implemented |
| Audit evidence package export | Generate ZIP package with evidence images, SHA-256 manifest, human-readable report, and package README. | Implemented |
| Print/PDF audit report | Generate a printable HTML report designed for browser Save-as-PDF. | Implemented |
| Expanded Controls | Replace Lockdowns with Controls, support multiple frameworks, status/profile/rationale, authorized imports, and official resource links. | Implemented |
| Framework support beyond CIS | Support CIS, NIST 800-53, NIST CSF, CMMC, SOC 2, ISO 27001, PCI DSS, HIPAA, STIG, and custom controls without redistributing licensed benchmark text. | Implemented |
| Findings/remediation tracker | Add `FND-####` records with severity, status, owner, due date, risk, recommendation, remediation, controls, and evidence. | Implemented |
| Note categories | Add note type and priority filters for audit requests, findings, questions, follow-ups, observations, and remediation. | Implemented |
| PBC/auditor request tracker | Add `PBC-####` records with owner, due date, status, auditor reference, notes, controls, and evidence mappings. | Implemented |
| AI workflow actions | Add one-click prompts for evidence review, evidence gaps, failed controls/findings, and audit summarization. | Implemented |
| Screenshot-aware AI | Require explicit opt-in and selected artifact before image pixels are sent; support Ollama vision and compatible multimodal APIs. | Implemented |
| AI evidence-gap analysis | Include structured workspace context so the assistant can identify controls/requests with missing or weakly documented evidence. | Implemented |
| Workspace encryption | Optional AES-GCM encrypted IndexedDB state with PBKDF2-SHA-256 password-derived key. | Implemented |
| External AI data warnings / credential controls | Make screenshot transmission explicit, distinguish connectors, exclude credentials from backups by default, and provide Clear Credentials. | Implemented |
| Home audit-readiness dashboard | Show evidence/control/PBC/finding KPIs, progress, attention items, recent evidence, and recent activity. | Implemented |
| Activity log | Record key workspace actions in a view-only trail with CSV export. | Implemented |
| Duplicate detection | Compare SHA-256 hashes within the active workspace before accepting a screenshot. | Implemented |

## Navigation

### Home

The Home dashboard summarizes the active audit workspace, including:

- Evidence count and review progress
- Control completion
- PBC request completion
- Open and overdue findings
- Overall workflow-readiness percentage
- Items needing attention
- Recent evidence
- Recent activity

The readiness metric is a workflow-completion aid only. It is not an audit opinion and does not determine compliance.

### Artifacts

Artifacts are screenshot evidence. Each accepted image receives:

- Sequential evidence ID such as `ART-0001`
- Upload timestamp
- Original filename
- File size
- MIME type
- SHA-256 integrity hash
- Evidence workflow status
- Owner
- System/application
- Collection method
- Description of what the evidence proves
- Internal note
- Auditor comment/response context
- Tags
- Related controls
- Related PBC requests

Supported browser-readable image types include PNG, JPG/JPEG, WEBP, and GIF. Files larger than 25 MB are rejected.

The app calculates SHA-256 before accepting an artifact. If an identical image already exists in the active audit workspace, the new upload is rejected as a duplicate and the existing evidence ID is shown.

Evidence workflow statuses are:

- Collected
- Reviewed
- Accepted
- Needs Update
- Rejected

Artifacts can be searched and filtered by status, mapped control, and tags.

#### Annotation and redaction

From an evidence record, choose **Annotate / redact** to create a new evidence version. Available tools are:

- Highlight rectangle
- Redact rectangle
- Arrow
- Undo

Saving the annotation creates a new artifact with its own evidence ID, timestamp, hash, and a reference to the source evidence. The original evidence is preserved.

### Notes

Notes are workspace-level audit notes. They support the following types:

- General
- Auditor Request
- Finding
- Question
- Follow-up
- Observation
- Remediation

Priority can be High, Medium, or Low. Notes are filterable by type and priority.

### Controls

The former **Lockdowns** area is now **Controls**.

Controls can represent CIS, NIST 800-53, NIST CSF, CMMC, SOC 2, ISO 27001, PCI DSS, HIPAA, STIG, or organization-specific requirements. Each control supports:

- Framework
- Control ID
- Title
- Recommendation/requirement summary
- Rationale/audit note
- Status
- Profile
- Evidence mappings
- Finding mappings

Control statuses are:

- Review
- Pass
- Fail
- N/A
- Needs Evidence

The app ships only a small generic CIS-oriented starter checklist. It does **not** redistribute licensed CIS Benchmark text.

Authorized control content can be imported from JSON or CSV. An authorized HTTP/HTTPS JSON/text endpoint can also be used when that endpoint permits browser CORS requests.

Recommended official resources are linked from the Controls page:

- CIS Benchmarks
- CIS Controls
- NIST SP 800 publications
- NIST Cybersecurity Framework

### PBC Requests

The **PBC Requests** page tracks Provided By Client / auditor request-list items. Each request receives an ID such as `PBC-0001` and supports:

- Request text
- Owner
- Due date
- Status
- Auditor reference
- Notes
- Related controls
- Mapped evidence

Statuses are:

- Not Started
- Collecting
- Ready
- Submitted
- Accepted

Evidence and PBC mappings are kept bidirectionally consistent by the application.

### Findings

Findings receive IDs such as `FND-0001` and support:

- Finding title
- Severity
- Status
- Owner
- Due date
- Risk/observation
- Recommendation
- Remediation note
- Related controls
- Evidence/remediation evidence

Severity values are High, Medium, Low, and Informational.

Statuses are:

- Open
- In Progress
- Remediated
- Accepted Risk
- Closed

### Chat

Chat is constrained to audit-related work, including:

- Auditing
- Evidence/screenshots
- Controls and frameworks
- PBC requests
- Findings and remediation
- Audit notes
- App settings
- Security configuration and hardening

The app includes shortcut prompts for evidence review, evidence-gap analysis, failed-control/finding review, and audit summarization.

#### Ollama

Ollama can be local or remote.

Examples:

- `http://localhost:11434`
- `http://192.168.1.50:11434`
- A protected internal HTTPS reverse-proxy URL

Features include:

- Configurable Ollama server URL
- Model discovery through `/api/tags`
- Optional Bearer token for authenticated proxies
- Connection testing
- Chat through `/api/chat`

A remote Ollama server or reverse proxy must allow CORS requests from the Audit Prep Tool origin. With the packaged launcher, the browser origin is `http://127.0.0.1:51327`.

Do not expose an unauthenticated Ollama service directly to the public Internet.

#### OpenAI-compatible APIs

The application also supports a configurable OpenAI-compatible chat-completions endpoint with:

- Full endpoint URL
- Model name
- API key
- Live connection test

Compatibility depends on the provider implementing the expected chat-completions request and response shape.

#### Screenshot-aware AI

Screenshot pixels are **not** sent by default.

To send an image, the user must explicitly:

1. Enable **Allow selected screenshot pixels to be sent to AI**.
2. Select a specific artifact.
3. Send a chat message.

For Ollama, the image is attached using Ollama's image-message format. For OpenAI-compatible providers, the app uses the common `image_url` multimodal message format. The selected model/provider must support vision input.

Without that explicit option, AI receives only text workspace context such as evidence metadata, notes, controls, requests, and findings.

### Activity

Important workspace actions are recorded in a view-only activity trail, including actions such as:

- Evidence uploads, edits, annotations, and deletion
- Note creation and updates
- Control import/edit/delete operations
- PBC request changes
- Finding changes
- AI use
- Backup/package export
- Workspace/security changes

The activity trail can be searched and exported as CSV. The UI provides no direct edit/delete function for individual activity entries.

### Reset / Data management

The Reset page contains backup, import, reporting, workspace protection, credential handling, and destructive reset functions.

#### Backup export

A JSON backup contains all audit workspaces and embedded screenshot data. API keys and Ollama Bearer tokens are excluded by default.

Select **Include API keys / bearer tokens** only when a credential-inclusive backup is specifically required and will be protected appropriately.

#### Import

Importing a v3 backup replaces the current local application state after confirmation.

#### Printable report / PDF

**Printable report** creates an audit report in a browser print view with:

- Workspace metadata
- Executive summary
- Control status table
- PBC request table
- Finding table
- Evidence appendix with screenshots and hashes

Use the browser's **Save as PDF** feature when a PDF is needed.

#### Audit package ZIP

**Audit Package ZIP** creates a standard uncompressed ZIP containing:

- `Evidence-Manifest.csv`
- `Audit-Report.html`
- `README.txt`
- An `Evidence/` directory containing each screenshot named with its evidence ID

The evidence manifest includes SHA-256 hashes so recipients can verify evidence integrity.

## Multiple audit workspaces

Use **Audits** in the top bar to create and switch audit engagements. Workspace metadata includes:

- Audit name
- Organization/system
- Auditor/assessor
- Audit lifecycle status
- Primary framework
- Benchmark/framework version
- Start date
- Target date

Deleting an audit workspace removes its contained evidence and related records from local application state after confirmation.

## Global search

The search box in the top bar searches the active workspace across evidence, controls, PBC requests, and findings. Press **Enter** to view matching records.

## Themes

The application includes six built-in themes available from every page:

1. Dark
2. Light
3. Slate
4. Forest
5. Sunset
6. Ocean

The selected theme persists locally.

## Storage and privacy

Application state is stored in browser IndexedDB for the Audit Prep Tool origin. The fixed launcher address is intentional so the same browser origin is used across launches.

Data remains local unless a user:

- Exports it
- Downloads an audit package/report
- Configures and uses an AI service
- Explicitly attaches screenshot pixels to an AI request

Browser-profile deletion, browser storage cleanup, different browser profiles, or enterprise browser policy can affect local data availability. Regular backup exports are recommended.

## Optional local storage encryption

Audit Prep Tool can optionally encrypt application state at rest inside IndexedDB.

Implementation:

- AES-GCM encryption
- Password-derived key using PBKDF2-SHA-256
- 250,000 PBKDF2 iterations
- Random salt
- Random AES-GCM IV for each save
- Password is not stored by the application

The encryption password is required when the app is reopened. If the password is lost, the encrypted local workspace cannot be recovered by the application. Keep a protected backup according to organizational policy.

This feature protects the app's stored state from casual inspection of the IndexedDB record. It does not replace operating-system full-disk encryption, endpoint controls, browser-profile protections, or enterprise key-management requirements.

## Credentials

AI credentials are stored in the application state for the current browser profile unless cleared.

Use **Clear credentials** on the Chat page to remove the stored OpenAI-compatible API key and Ollama Bearer token without deleting the rest of the workspace.

Credential values are omitted from backup exports by default.

## Security considerations

Audit evidence can contain sensitive data. Recommended operational practices include:

- Run the supplied launcher only on trusted endpoints.
- Use full-disk encryption where required by policy.
- Redact secrets, personal data, tokens, or unnecessary identifiers from screenshots.
- Prefer protected internal Ollama endpoints for sensitive AI use.
- Review evidence before sending screenshot pixels to any external AI service.
- Protect exported JSON backups and audit packages according to their data classification.
- Code-sign launchers before broad enterprise deployment.
- Do not use application workflow statuses as a substitute for auditor judgment.

## Technical architecture

The UI is a single self-contained `index.html` file containing HTML, CSS, and JavaScript.

`launcher.go` uses only the Go standard library and:

- Embeds `index.html` at compile time
- Binds only to `127.0.0.1:51327`
- Serves the embedded UI
- Exposes a small local health endpoint
- Opens the default browser

End users do not need Go. Go is required only when rebuilding the launcher binaries from source.

## Development and source control

The maintained source of record is the repository-root `index.html` and `launcher.go`. The UI remains intentionally single-file so the launcher can embed it with the Go standard library. Packaged binaries, checksum manifests, and release-test output are release artifacts and are not committed to the normal source branch; build them from source for a release.

Useful local checks:

```powershell
node tests/regression-checks.mjs
go build -o build/audit-prep-tool.exe launcher.go
```

See [architecture documentation](docs/architecture.md), [security notes](docs/security.md), [testing guidance](docs/testing.md), and the [release process](docs/release-process.md). Git history before v3.0.2 was reconstructed from the existing application, release notes, and development records during migration into source control; it is not presented as the original commit history.

## Rebuilding launchers

From the directory containing `launcher.go` and `index.html`:

```bash
GOOS=windows GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o audit-prep-tool-windows-amd64.exe launcher.go
GOOS=linux   GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o audit-prep-tool-linux-amd64 launcher.go
GOOS=darwin  GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o audit-prep-tool-macos-amd64 launcher.go
GOOS=darwin  GOARCH=arm64 go build -trimpath -ldflags="-s -w" -o audit-prep-tool-macos-arm64 launcher.go
```

Regenerate `SHA256SUMS.txt` after rebuilding.

## Data compatibility

Version 3 introduces a workspace-oriented data model. The application performs best-effort migration of earlier local data where available. A Version 3 JSON backup uses the current multi-workspace format.

Because Version 3 materially changes the data model, create and retain a backup before replacing an older deployed version.

## Current limitations

- The application is single-user and browser-profile local; it is not a multi-user collaboration server.
- Activity logging is application-level and is not a cryptographically signed append-only ledger.
- Framework content is not automatically licensed or redistributed by the app; organizations remain responsible for authorized benchmark content.
- Cross-origin control imports and remote AI calls depend on target-server CORS policy.
- Vision support depends on the configured AI model and provider.
- Browser storage quotas vary by browser and enterprise policy.
- The app does not replace a GRC platform, formal evidence repository, enterprise identity system, or auditor judgment.

# Changelog

All future product changes should be recorded in this section when the application is modified.

## 3.0.2 - 2026-09-18

### User testing and integration fix

- Completed click-level user testing across navigation, all six themes, workspaces, evidence upload/validation, evidence metadata and filters, Controls, PBC requests, Findings, Notes, dashboard metrics, activity tracking, control imports, remote control fetches, report/package exports, backup/import, workspace encryption controls, screenshot annotation/redaction, and both Ollama and OpenAI-compatible AI workflows.
- Verified remote Ollama URL handling, optional Bearer authentication, model discovery, audit-only chat gating, explicit screenshot-vision opt-in, and OpenAI-compatible multimodal payloads.
- Fixed annotation/PBC relationship consistency: annotated or redacted evidence versions now synchronize inherited PBC request mappings in both directions immediately after creation.
- Re-tested control deletion, PBC deletion, duplicate detection, 25 MB upload rejection, ZIP integrity, credential-excluded backups, and activity CSV export.
- Rebuilt all self-contained desktop launchers from the 3.0.2 UI and regenerated release checksums.

## 3.0.1 - 2026-09-18

Integration and relationship-consistency release.

- Added delete actions for existing PBC requests and findings with confirmation prompts.
- Deleting a PBC request now removes its reverse mappings from evidence.
- Deleting a control now clears the control from evidence, findings, and PBC requests.
- Corrected activity logging so newly created controls are recorded as created rather than saved.
- Added the recommendation-by-recommendation implementation plan and status matrix to this README.
- Rebuilt all platform launchers from the 3.0.1 embedded UI and regenerated release checksums.

## 3.0.0 - 2026-09-18

Major audit-workflow release.

### Navigation and product structure

- Renamed **Lockdowns** to **Controls**.
- Added PBC Requests, Findings, and Activity navigation areas.
- Added multi-audit workspace support and workspace selector.
- Added workspace metadata for organization/system, assessor, lifecycle status, primary framework, benchmark/framework version, and audit dates.
- Added six persistent UI themes across all pages.

### Dashboard

- Rebuilt Home as an audit-readiness dashboard.
- Added evidence, control, PBC, findings, and readiness KPI cards.
- Added workflow progress indicators.
- Added items-needing-attention detection.
- Added recent-evidence and recent-activity panels.

### Evidence / Artifacts

- Added sequential evidence IDs.
- Added upload timestamps and structured evidence metadata.
- Added SHA-256 integrity hashing.
- Added duplicate screenshot detection using SHA-256.
- Added review workflow statuses.
- Added tags, search, status filtering, control filtering, and tag filtering.
- Added evidence owner, system/application, collection method, evidence description, internal note, and auditor comment fields.
- Added control and PBC request mapping.
- Added annotation, highlighting, arrows, and blackout redaction.
- Annotation/redaction creates a new evidence version instead of overwriting the source.
- Added direct evidence image download.

### Controls

- Added generic framework support for CIS, NIST, CMMC, SOC 2, ISO 27001, PCI DSS, HIPAA, STIG, and custom controls.
- Added control IDs, profiles, recommendations, rationale, status, evidence count, and open-finding count.
- Added authorized JSON/CSV import.
- Added authorized HTTP/HTTPS JSON/text endpoint import.
- Added links to official CIS and NIST resources.
- Preserved the policy of not embedding licensed CIS Benchmark text.

### PBC requests

- Added sequential PBC IDs.
- Added request owner, due date, status, auditor reference, notes, related controls, and evidence mappings.
- Added bidirectional PBC/evidence mapping.

### Findings

- Added sequential finding IDs.
- Added severity, owner, due date, workflow status, risk, recommendation, remediation notes, related controls, and evidence mappings.

### Notes

- Added note types and priorities.
- Added note filtering.

### AI

- Preserved local and remote Ollama support.
- Preserved optional Bearer-token support for protected Ollama proxies.
- Preserved OpenAI-compatible API support.
- Added one-click audit prompt actions.
- Added richer workspace context for controls, requests, findings, notes, and evidence metadata.
- Added explicit opt-in screenshot vision support for Ollama and compatible multimodal APIs.
- Screenshot pixels remain off by default.
- Added clear-credentials action.

### Reporting and export

- Added evidence manifest CSV export.
- Added activity CSV export.
- Added printable audit report suitable for browser Save-as-PDF.
- Added auditor-ready ZIP package containing report, manifest, README, and evidence images.
- Preserved full JSON backup/import.
- Credentials remain excluded from backups by default.

### Security and integrity

- Added optional AES-GCM application-state encryption at rest.
- Added PBKDF2-SHA-256 password-based key derivation.
- Added password change and encryption-disable controls.
- Added application activity trail.
- Added external AI data-handling notices.
- Added evidence SHA-256 hashes to reports/manifests.

### Packaging

- Retained self-contained Windows, Linux, Intel macOS, and Apple Silicon macOS launchers.
- Retained loopback-only local serving on `127.0.0.1:51327`.
- Retained standalone `index.html` fallback.

## 2.x beta line

- Added local/remote Ollama connectivity and model discovery.
- Added OpenAI-compatible API connection testing.
- Added optional Ollama Bearer token support.
- Fixed stale Chat connection settings.
- Added upload validation and 25 MB screenshot limit.
- Excluded credentials from exports by default.
- Added the original self-contained Go launchers and beta test documentation.

---

When changing the product in future releases, update **Version** at the top and append a dated entry under **Changelog** describing user-visible changes, security changes, migration considerations, and packaging changes.
