# `internal/` — package layout (work in progress)

This tree is being filled in over a multi-stage migration that moves pike's
flat `src/` package into a proper hierarchy. While the migration is in
flight, `src/` continues to compile and run; `internal/` packages are
populated stage by stage and `src/` gradually thins out into re-export
shims.

If you wandered in here mid-migration and aren't sure which file owns what,
the answer right now is "almost certainly still in `../src/`". Anything
under `internal/` is the new home for code with a clearer boundary.

## Target tree

```
internal/
├── policy/        # pure types + assembly: Sorted, Policy, Statement, OutputPolicy,
│                  #   PolicyDiff, RuntimePermission, IAMBinding,
│                  #   GetPolicy, Unique, Minify, NewStatement, WriteOutput
│                  # (no I/O, no provider knowledge)
├── iac/           # Terraform/OpenTofu HCL parsing + module resolution:
│                  #   ResourceV2, GetResources, GetBlockAttributes,
│                  #   DetectBackend, ExtractIAMBindings, modules.go contents,
│                  #   Init/LocateTerraform/GetTF (tfexec wrappers)
├── provider/
│   ├── aws/       # AWS-specific permissions, datasource, files (embed),
│   │              #   policy templating, mapping/ JSON
│   ├── gcp/       # GCP equivalents + project discovery
│   └── azure/     # Azure equivalents
├── iam/           # talks to deployed IAM (not local code):
│                  #   Watch, Compare (AWS + GCP), Inspect
├── credentials/   # AWS STS credential acquisition (setAWSAuth/unSetAWSAuth)
├── github/        # GitHub API: dispatch (Invoke), secrets (Remote)
├── git/           # go-git clone wrapping (Pull/Repository)
├── scan/          # Scan + Runtime + MakePolicy orchestration
├── make/          # Make + Apply (terraform-exec wrapping)
├── readme/        # Readme + ReplaceSection
└── util/          # truly generic helpers: FileExists, AlmostEqual, etc.
```

`cmd/pike/main.go` will import these directly; there is no `internal/cli/`
adapter layer.

`parse/` (the resource extractor used by the weekly resources.yml workflow)
lives at the repo root, not under `internal/`, because it is invoked by
external contributors during mapping work and benefits from being
discoverable.

## Migration stage progress

- [x] Stage 0 — CI safety net (golden help diff + API surface snapshot).
- [ ] Stage 1 — `git mv src/parse parse`.
- [ ] Stage 2 — `main.go` → `cmd/pike/main.go` + Dockerfile/goreleaser/Makefile updates.
- [ ] Stage 3 — `internal/policy/` extraction; `src/` re-exports via type aliases.
- [ ] Stage 4 — `internal/iac/` extraction.
- [ ] Stage 5 — all three `internal/provider/` packages in one PR.
- [ ] Stage 6 — credentials, github, git, iam, scan, make, readme, util.
- [ ] Stage 7 — delete `src/`, wire `cmd/pike/main.go` through `internal/` directly.

## Before deleting `src/` (Stage 7 pre-flight)

Confirm nobody is importing `github.com/jameswoolfenden/pike/src` as a
library by checking deps.dev:

```sh
curl -s "https://api.deps.dev/v3alpha/systems/go/packages/github.com%2Fjameswoolfenden%2Fpike/dependents" | jq
```

The grep on this repo at the start of Stage 0 found only internal callers
(main.go + own tests + own coverage tool); the deps.dev check is the
external safety net.
