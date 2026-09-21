# Security and privacy

## Implemented protections

- Evidence is stored in browser IndexedDB for the fixed local application origin; it is not uploaded by the launcher.
- Each accepted artifact receives a browser-calculated SHA-256 hash. The hash supports integrity checking but is not a signed chain of custody.
- Duplicate detection compares SHA-256 values within the active workspace.
- API keys and Ollama bearer tokens are excluded from JSON backups unless the user explicitly selects credential inclusion. The UI can clear those credentials separately.
- Screenshot pixels are not sent to an AI provider by default. The user must enable vision sharing and select a particular artifact.
- Optional app-managed encryption uses AES-GCM and a PBKDF2-SHA-256-derived key. Lost passwords cannot be recovered by the app.
- The launcher binds only to `127.0.0.1:51327` and does not expose an HTTP listener on the LAN.

## Operational considerations

Browser-profile cleanup, profile changes, and storage quotas can affect local data. Regular protected backups are recommended. App-managed encryption does not protect evidence displayed in the browser, prevent endpoint compromise, or replace full-disk encryption and organizational access controls.

AI calls can disclose their supplied text context to the configured provider. Vision calls can disclose selected screenshot pixels. Prefer protected internal Ollama deployments for sensitive material; do not expose an unauthenticated Ollama service publicly. Remote services must permit the application's origin through their CORS policy.

Audit packages, reports, and JSON backups can contain sensitive evidence. Handle them according to their data classification. The activity trail is application-level and is not a cryptographically signed, append-only ledger. Workflow statuses and AI output are not auditor conclusions.
