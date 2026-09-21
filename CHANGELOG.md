# Changelog

All notable changes to Audit Prep Tool are documented here.

Git history before v3.0.2 was reconstructed from the existing application, release notes, and development records when the project was migrated into source control. Entries marked as reconstructed describe the known release state; they do not claim exact historical commit timing.

## [3.0.2] - 2026-09-18

### Fixed

- Annotated or redacted evidence now synchronizes inherited PBC mappings back to the linked PBC request immediately.

### Verified

- Completed regression coverage of the audit workflow, exports, encryption controls, and AI connector privacy boundaries.
- Rebuilt the four desktop launchers and regenerated checksums for the release.

## [3.0.1] - 2026-09-18

### Fixed

- Deleting a control clears its relationships from evidence, PBC requests, and findings.
- Deleting a PBC request clears its reverse evidence relationships.
- Deleting a finding preserves its related evidence.
- New-control activity entries are recorded as creations.

### Documentation

- Added implementation-status documentation to the product README.

## [3.0.0] - Reconstructed history

### Added

- Multi-workspace audit model, evidence IDs and hashes, duplicate detection, review statuses, tags, mappings, global search, dashboard metrics, PBC tracking, findings, activity log, and exports.
- Framework-neutral controls, optional encrypted IndexedDB storage, annotation/redaction evidence versions, and controlled AI workflows.
- Audit package ZIP and printable audit report generation.

## Earlier development - Reconstructed history

- Initial local-first screenshot, notes, themes, and browser-storage application.
- Configurable local and remote Ollama support.
- Upload validation, OpenAI-compatible connector testing, credential-excluded backups, and audit-only AI scope controls.
- Self-contained desktop launchers for Windows, Linux, and macOS.
