# Security Policy

## Reporting a Vulnerability

**Please do not open a public GitHub issue for security vulnerabilities.**

Report privately using GitHub's Security Advisory workflow:

👉 **[Report a vulnerability](https://github.com/JamesWoolfenden/pike/security/advisories/new)**

This opens a private channel between you and the maintainers where we can
discuss, reproduce, and patch the issue before disclosure.

When reporting, please include:

- A description of the vulnerability and its impact
- Steps to reproduce (a minimal Terraform/OpenTofu snippet is ideal)
- The version of `pike`, OpenTofu/Terraform, and the affected provider
- Any suggested mitigation or fix

## Supported Versions

Only the **latest released minor version** is supported. Security fixes are
released on top of the current minor; older minors do not receive backports.

| Version | Supported |
| ------- | --------- |
| latest  | ✅ |
| older   | ❌ |

## Response Targets

These are targets, not guarantees — this is a solo-maintained OSS project.

| Severity (CVSS) | Acknowledgement | Fix / mitigation |
| --------------- | --------------- | ---------------- |
| Critical (9.0+) | 2 business days | 14 days |
| High (7.0–8.9)  | 3 business days | 30 days |
| Medium (4.0–6.9)| 5 business days | 60 days |
| Low (<4.0)      | 10 business days| 90 days |

## Disclosure

Once a fix is available we will:

1. Release a patched version.
2. Publish a GitHub Security Advisory (and request a CVE if appropriate).
3. Credit the reporter in the advisory, unless anonymity is requested.

## Scope

In scope:

- The `pike` CLI binary and its Go source under this repository.
- The official Docker image published from this repository.
- Published release artefacts (Homebrew tap, Scoop bucket, GitHub Releases).

Out of scope:

- Third-party forks or unofficial packages.
- Vulnerabilities in upstream dependencies — please report those to the
  upstream project. If the upstream fix requires a pinned bump here, we're
  happy to ship it.
- Permissions output that is over-permissive for a specific deployment: the
  tool's output is a starting point (see the "Limitations" language in
  the README). A policy being more permissive than your risk tolerance is a
  configuration concern, not a vulnerability.
