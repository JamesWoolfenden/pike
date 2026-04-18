# Pike — Code Review

_A thorough code review of the master branch as of 2026-04-17. Findings are grouped by severity with specific file/line references and concrete fixes. Every major claim was verified against the actual tree, not just inferred from naming._

---

## Overall take

Pike is a mature, well-scoped OSS project with real teeth: OpenSSF badge, cross-platform CI matrix, pre-commit hooks, codecov, Homebrew + Scoop + Docker distribution, and 29k lines of Go covering three cloud providers. The bones are good.

The gaps are the ones you'd expect in a tool that started as a personal project and grew past that: one flat `src/` package that's outgrown a single directory, three logging frameworks living side-by-side, a release pipeline stitched together from PowerShell + bash + YAML, and CI that runs tests but doesn't run the lint / security / vuln-scan tools you already have in your `Makefile`. None of this is fatal. All of it is the distance between "works" and "best of me."

The one finding that would genuinely embarrass a stranger reading the code for the first time is a hardcoded personal AWS account ID baked in as a CLI default. Fix that first.

---

## Critical — fix before next release

### C1. Hardcoded personal AWS account ID in CLI defaults
**`main.go:342`** — the `watch` subcommand has `Value: "arn:aws:iam::680235478471:policy/basic"` as the default ARN. That account ID is traceable to you, the default is never useful to anyone else, and it quietly advertises an internal account number in every `pike --help` output.

**Fix:** delete the `Value` line. `arn` should be a required flag or, at most, default to empty and error out clearly with a message pointing users to `--arn` / `$ARN`.

### C2. `panic()` in `GetPolicy` instead of returning an error
**`src/policy.go:204`** — `tmpl.Execute` returns an error, and the code calls `panic(err)` instead of returning a `&templateExecuteError{err}` the way the parse-error path does three lines above. A malformed template (or nil data) will tear down the CLI mid-run.

**Fix:**
```go
if err := tmpl.Execute(&output, theDetails); err != nil {
    return OutPolicy, &templateExecuteError{err}
}
```
Then add the error type to `src/error.go` alongside the existing `templateParseError`.

### C3. CI runs `go test` but nothing else — your own Makefile has lint, gosec, govulncheck
**`.github/workflows/ci.yml`** and **`pr.yml`** run only `go build` + `go test`. The `Makefile` defines `lint`, `gosec`, `staticcheck`, `govulncheck`, and `install-tools` targets — none of them gate merges. That means the entire supply-chain / static-analysis story lives only on your laptop.

**Fix:** add to `pr.yml` after the existing test step:
```yaml
- uses: golangci/golangci-lint-action@v6
  with: { version: v1.64 }
- run: go run golang.org/x/vuln/cmd/govulncheck@latest ./...
- run: go install github.com/securego/gosec/v2/cmd/gosec@latest && gosec -quiet ./...
```
Also create `.golangci.yml` at the root pinning the linter set — right now `make lint` uses golangci-lint's defaults, which drift between versions.

### C4. PR workflow has a `docs` job that does nothing
**`.github/workflows/pr.yml:56-67`** — the `docs` job checks out the repo, installs Go, then stops. No lint, no build, no validation. It passes trivially and wastes a runner slot on every PR.

**Fix:** either delete the job, or give it a purpose (`markdownlint README.md CONTRIBUTING.md SECURITY.md` is a good starter).

### C5. Dockerfile downloads a prebuilt binary instead of building from source
**`Dockerfile`** — runs `alpine:3.18.2` → `apk add build-base git curl jq bash` → curls the latest GitHub release asset → untars. This means:
- `docker build .` against a specific git SHA doesn't produce a binary from that SHA. It produces whatever is currently tagged `latest`.
- Base image `alpine:3.18.2` was released 2023 and hasn't been bumped in 2+ years; it's accumulating CVEs.
- Container runs as root (no `USER` directive).
- `build-base` pulls in gcc/make just to `tar`.

**Fix:** switch to a multi-stage build:
```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.Version=${VERSION}" -o /pike

FROM alpine:3.20
RUN apk add --no-cache ca-certificates git && adduser -D -u 10001 pike
COPY --from=builder /pike /usr/bin/pike
USER pike
ENTRYPOINT ["/usr/bin/pike"]
```

### C6. SECURITY.md is nine lines with a plaintext email
**`SECURITY.md`** — the document that a security researcher will look at first directs them to a Gmail address, has no PGP key, no advisory channel, no supported-versions table, no SLA. GitHub has built-in private security advisories — enable them and point to them.

**Fix:** replace the file with something like:
```markdown
# Security Policy

## Reporting a vulnerability
Please report privately via
https://github.com/JamesWoolfenden/pike/security/advisories/new

Do NOT open a public issue for suspected vulnerabilities.

## Supported versions
Only the latest minor release is supported.

## Response targets
- Acknowledgement within 3 business days
- Fix or mitigation within 30 days for CVSS ≥ 7, 90 days otherwise
```

---

## Important — structural and consistency issues

### I1. One flat `src/` package is doing too much
`src/` has ~76 `.go` files in a single `package pike`, including three god-files:

| File | Lines | What's in it |
|------|-------|--------------|
| `src/files_gcp.go` | 3,338 | `//go:embed` wrappers for GCP mappings |
| `src/files.go` | 2,867 | Same, cross-provider |
| `src/aws.go` | ~1,800 | All AWS logic |

The `parse/`, `schema/`, and `mapping/` subtrees show you already understand the pattern. Promoting logical boundaries to packages would give you:
- `internal/provider/aws`, `internal/provider/gcp`, `internal/provider/azure` — cloud-specific logic and embeds co-located
- `internal/policy` — template rendering, compare, inspect
- `internal/iac` — HCL/Terraform parsing
- `cmd/pike` — CLI entry (today `main.go` is at root, which is fine, but the subcommand setup could move)

Do this incrementally — pick one provider, extract it, land the PR, repeat. The 3,338-line `files_gcp.go` collapses to a handful of files when it lives next to the GCP code it describes.

### I2. Three logging frameworks
The tree imports `github.com/rs/zerolog` **and** uses `log.Printf` (stdlib) **and** `fmt.Println` / `fmt.Print` for user-facing output. Examples: `src/data.go:368-387`, `src/credentials.go:69-77`, `src/compare.go:163,243`, `src/gitHub.go:76`.

This makes structured logging useless (you can't grep a log level that isn't there) and produces unpredictable routing (stdlib `log` goes to stderr by default; `fmt.Println` goes to stdout).

**Fix:** decide the rule:
- `zerolog` for internal diagnostics/telemetry
- a dedicated `ui` / `out` helper for user-facing output (uses stdout, respects `--quiet` / `--json`)

Then ban direct `fmt.Print*` and stdlib `log` anywhere outside that helper — a golangci-lint `forbidigo` rule enforces it.

### I3. Two error-construction patterns fighting each other
`src/error.go` defines ~50 typed errors (`templateParseError`, `assertionFailedError`, etc.), but much of the codebase uses `fmt.Errorf("op: %w", err)` instead. Example: `src/aws.go:1687` vs the typed-error pattern three functions later in the same file.

Pick one. The typed-error pattern is valuable when callers want to branch with `errors.As`, but most of your errors are wrapped once and logged — `fmt.Errorf` with `%w` is enough and removes ~300 lines of boilerplate.

### I4. Public functions return AWS SDK types
**`src/credentials.go:83`** returns `*sts.AssumeRoleOutput`. Any consumer (importer or test) now transitively depends on `aws-sdk-go-v2/service/sts` and is pinned to your SDK major version. When AWS SDK v3 ships, every caller breaks.

**Fix:** define a small internal struct and convert at the boundary:
```go
type AWSCredentials struct {
    AccessKeyID, SecretAccessKey, SessionToken string
    Expires time.Time
}
```

### I5. `os.Exit(1)` from inside a CLI `Action` handler
**`main.go:275-279`** (`compare` command) calls `os.Exit(1)` in the middle of the action closure. urfave/cli already handles exit codes when you return an error. Calling `os.Exit` bypasses deferred cleanup and makes the function untestable.

**Fix:** return the error. Let urfave/cli translate it to an exit code.

### I6. Ignored errors on the credential-setting path
**`src/credentials.go:100-102, 108-110`** use `_ = os.Setenv(...)`. If your credential-env setup silently fails on Windows (it can — UAC + certain env stores), the subsequent AWS call fails several stack frames away with a generic "no credentials" message. Hours of debugging.

**Fix:** join errors, don't drop them.

### I7. Unchecked type assertion on `sync.Map` value
**`src/scan.go:579`**: `mutex := dirMutex.(*sync.Mutex)`. If the map ever holds a different type (refactor, test, concurrent init bug), this panics in production.

**Fix:** comma-ok form and return an internal error. Same file has `initMutex.LoadOrStore(...)` ignored on line 578 — review the pair together.

### I8. Copy-paste docstring on `GetAZUREPermissions`
**`src/azure.go:68`** — the godoc reads "GetAZUREPermissions for GCP resources". Small, but it's the exact kind of thing a reader notices on first skim and starts doubting everything else.

### I9. `secrets_Internal_test.go` vs `secrets_test.go` — inconsistent casing
**`src/secrets_Internal_test.go`** (capital I) is the odd one out; every other internal-test file uses lowercase. Some filesystems are case-insensitive (Windows, macOS default), some aren't (Linux). Rename.

### I10. Go version skew between PR and master CI
`pr.yml:15` uses `go-version: 1.25.x` (floating). `ci.yml:19` uses `1.25.8` (pinned). A PR that passes 1.25.12 could regress in master on 1.25.8. Pin both to `1.25.8` or loosen both; don't mix.

### I11. No module caching in `ci.yml`
`pr.yml` has an `actions/cache` step for `~/go/pkg/mod`. `ci.yml` doesn't. Master builds download modules from scratch every push — a few minutes wasted per commit.

### I12. Commit-author email in `resources.yml` is your personal Duck address
**`.github/workflows/resources.yml`** sets `git config user.email "jim.wolf@duck.com"` for the bot commit. Use the GitHub Actions noreply address: `41898282+github-actions[bot]@users.noreply.github.com`. Keeps your personal email off public commit history and tells reviewers at a glance which commits are automation.

### I13. Incomplete TODOs on the `runtime` command
**`src/scan.go:179, 185`** say `// TODO: Implement AWS runtime permission formatting` and the Azure equivalent. The `runtime` subcommand is exposed in `main.go:198` but the formatter is stubbed. Either finish it, hide the subcommand behind a `--experimental` flag, or document the limitation in the help string.

### I14. README feature list doesn't match the actual CLI surface
README's "What's new" bullets mention "json modules support", the code has `src/modules.go` and `src/secrets.go` and `src/policy.go`, but the CLI subcommands are: `make`, `apply`, `remote`, `scan`, `runtime`, `compare`, `inspect`, `watch`, `readme`, `invoke`, `parse`, `pull`, `version`. No `modules`, `secrets`, or `policy` subcommand.

Either those capabilities are invoked via flags on `scan`/`compare` (document it), or they're internal-only (remove from README), or they're planned (put them behind a Roadmap section).

### I15. No `CHANGELOG.md`
Users upgrading from v1.x to v2.x have to read GitHub releases pages or `git log`. For an OSS tool that emits IAM policies — where a behaviour change can silently break a downstream Terraform plan — a proper changelog is not optional.

**Fix:** adopt Keep a Changelog format; `git-chglog` or `release-please` can generate it from conventional commits.

---

## Nice-to-have — polish

### N1. Split supporting docs out of the repo root
`PARSE_IMPROVEMENT.md`, `RELEASE_SUMMARY.md`, `RESOURCE_SCRIPT_README.md` all live at the top level. Move them to `docs/`. The repo-root `ls` should read as a clean, newcomer-friendly manifest.

### N2. No `CODEOWNERS`, `dependabot.yml`, `renovate.json`, PR template, Code of Conduct
Standard OSS hygiene. `dependabot.yml` in particular — you already have an OpenSSF badge; enabling Dependabot for `gomod` and `github-actions` is a five-line file and a weekly chore the bot handles for you.

### N3. Issue templates are unmodified GitHub defaults
`.github/ISSUE_TEMPLATE/bug_report.md` asks for "Desktop OS" and "browser", which is irrelevant for a CLI that only runs on a shell. Replace with fields that actually help you triage: `pike version`, `tofu/terraform version`, provider (AWS/Azure/GCP), a minimal `.tf` repro, full command + output.

### N4. `pike.jfif` is an odd logo format
`.jfif` is a legacy JPEG extension that some markdown renderers (and GitHub's image proxy, intermittently) don't handle. Re-export as `pike.png` and update README. Also: the current image is 92×92 — not bad, but a 200×200 or 400×400 would render crisper on high-DPI displays.

### N5. Release process is fragmented across languages
Version bumps involve `set-version.sh` (bash), `bump.ps1` (PowerShell), `release.ps1` (PowerShell), `.goreleaser.yml`, and `release.yml`. That's three languages for one job. Consolidate — GoReleaser's `before.hooks` can invoke a single script, and you can drop the PowerShell layer entirely if your dev OS supports WSL or Git Bash.

### N6. Release pipeline doesn't produce SBOM or sigstore attestation
`.goreleaser.yml` signs with GPG but doesn't emit an SBOM (`syft` / goreleaser's `sboms` directive) or a cosign signature. The OpenSSF badge rewards supply-chain provenance — picking this up is a high-signal, low-cost upgrade. Two lines in `.goreleaser.yml`:
```yaml
sboms:
  - artifacts: archive
```
and a `cosign sign-blob` step in the workflow.

### N7. `.orig` files in the working tree (but gitignored)
`src/files_gcp.go.orig` and `src/gcp.go.orig` exist on disk. Already in `.gitignore`, so they're not committed — but they've been sitting in your working tree since a merge conflict. `git clean -fd` or just delete them. Leaving them around is a minor code-review hazard if someone opens them thinking they're source.

### N8. 412 files in `git status` on master
Your working tree has 412 modified files (mostly `src/mapping/aws/**` + `terraform/aws/backup/**`) uncommitted on `master`. Not a code review finding, but worth noting: either commit the in-progress mapping update or branch/stash it. A contributor cloning and running `git pull` onto this state would be confused.

### N9. No "Limitations" section in README
You have the caveats language in the intro ("first step", "wildcarded", "no conditions yet"), but it's scattered. Promote it to a dedicated `## Limitations` section. Users who deploy a wildcard policy because they missed the caveat text are going to blame Pike, not themselves.

### N10. `tests` and `testacc` in Makefile but `testacc` is never referenced in CI
`make testacc` runs with `TF_ACC=1` and a 120m timeout — clearly your acceptance suite. It's not wired into any CI workflow and doesn't appear to have a dedicated lane (e.g., a nightly workflow). Either run it on a schedule, or drop it from the Makefile so it doesn't rot.

### N11. Codecov badge in README but no codecov.yml
`README.md:13` shows the codecov badge; the ci workflow uploads to Codecov; but there's no `codecov.yml` setting a target/threshold. Without a target, coverage regressions don't fail PRs. Minimal starter:
```yaml
coverage:
  status:
    project:
      default:
        target: auto
        threshold: 1%
```

### N12. `fmt.Errorf` without context vs with context
Some `fmt.Errorf("%w", err)` calls forward the error with no operation context, defeating the whole point. E.g. `src/aws.go:1748`. Pattern should always be `fmt.Errorf("operationName: %w", err)`.

### N13. Pre-commit runs `go test ./...` on every commit
`.pre-commit-config.yaml` includes `go-test` with `args: ["./..."]`. Given your testdata is 14 GB of locally-generated terraform init caches, a full `go test` on commit is not fast. Consider `-short` + dedicated skip tags on slow tests.

---

## Quick-win checklist (things you can knock out in an evening)

- [ ] Delete hardcoded ARN default in `main.go:342`
- [ ] Replace `panic(err)` in `src/policy.go:204` with typed error
- [ ] Rename `src/secrets_Internal_test.go` → `src/secrets_internal_test.go`
- [ ] Fix docstring typo in `src/azure.go:68`
- [ ] Delete `src/*.go.orig` files locally (already gitignored)
- [ ] Delete the empty `docs` job in `.github/workflows/pr.yml`
- [ ] Pin `pr.yml` go-version to `1.25.8` to match `ci.yml`
- [ ] Add module cache step to `ci.yml`
- [ ] Replace `.github/workflows/resources.yml` commit email with GitHub Actions noreply
- [ ] Expand `SECURITY.md` to point at GitHub private advisories
- [ ] Add `.github/dependabot.yml` for `gomod` + `github-actions`
- [ ] Add a `.golangci.yml` pinning the linter set
- [ ] Add `golangci-lint` + `govulncheck` steps to `pr.yml`

## Bigger refactors worth planning

1. **Package split** — break `src/` into `internal/provider/{aws,gcp,azure}` + `internal/policy` + `internal/iac` + `cmd/pike`. Do it provider-by-provider.
2. **Logging unification** — settle on zerolog internally + a small `ui` helper for user-facing output; enforce with `forbidigo`.
3. **Error model** — pick typed-errors OR `fmt.Errorf("%w")` and delete the other.
4. **Dockerfile** — go to multi-stage, non-root, hermetic build.
5. **Release pipeline** — collapse PowerShell + bash + YAML into a single GoReleaser-driven flow with SBOM + cosign.

---

_Reviewed against master @ ea40d9d5 on 2026-04-17._
