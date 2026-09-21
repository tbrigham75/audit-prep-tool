import { readFileSync } from 'node:fs';
import assert from 'node:assert/strict';

const ui = readFileSync(new URL('../index.html', import.meta.url), 'utf8');
const launcher = readFileSync(new URL('../launcher.go', import.meta.url), 'utf8');

const expected = [
  ["const APP_VERSION='3.0.2'", 'visible release version'],
  ["const DB_NAME='audit-prep-tool'", 'IndexedDB storage'],
  ["crypto.subtle.digest('SHA-256'", 'evidence hashing'],
  ["if(f.size>25*1024*1024)", 'upload size limit'],
  ['duplicate of ${dupe.evidenceId}', 'duplicate detection'],
  ["name:'AES-GCM'", 'encrypted state'],
  ['iterations:250000', 'PBKDF2 work factor'],
  ['syncArtifactLinks(a)', 'annotated-evidence PBC synchronization'],
  ['w.requests=w.requests.filter', 'PBC deletion path'],
  ['w.findings=w.findings.filter', 'finding deletion path'],
  ["$('allowVision').checked", 'explicit vision gate'],
  ["if(!$('includeSecrets').checked)", 'credential-excluded backup default'],
];

for (const [needle, label] of expected) {
  assert.ok(ui.includes(needle), `Missing ${label}`);
}
assert.ok(launcher.includes('const addr = "127.0.0.1:51327"'), 'Launcher must bind loopback only');
assert.ok(launcher.includes('mux.HandleFunc("/__health"'), 'Launcher must expose health endpoint');
console.log(`PASS: ${expected.length + 2} source-level regression checks`);
