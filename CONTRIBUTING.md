# Contributing to Pike

## Getting Started

1. Fork the repository and clone your fork.
2. Create a feature branch from `master`.
3. Make your changes, ensuring all checks pass (see below).
4. Open a pull request against `master`.

## Coding Standards

All contributions must meet the following requirements before they will be accepted:

### Formatting

Code must be formatted with `gofmt` and `goimports`:

```bash
gofmt -w ./...
goimports -w ./...
```

### Linting

Code must pass `golangci-lint` with the project configuration (`.golangci.yml`):

```bash
golangci-lint run
```

Enabled linters: `errcheck`, `forbidigo`, `govet`, `ineffassign`, `misspell`, `staticcheck`, `unused`.

Use `zerolog` for logging — do not use `log.Print`, `log.Printf`, or `log.Println`.

### Security

Code must pass `gosec` with no new findings:

```bash
make gosec
```

### Tests

All new functionality must include tests. Existing tests must continue to pass:

```bash
make test
```

If you change the CLI surface or exported package API, regenerate the golden files:

```bash
make goldens
```

### Adding Resource Mappings

New resource-type permission mappings live under `src/mapping/<provider>/resource/<service>/<type>.json`
(or `data/` for data sources). Filename must exactly match the Terraform resource type — pike resolves
mappings by filename, not by any registered list.

Every new mapping must also include a minimal example `.tf` fixture under
`terraform/<provider>/backup/<type>.tf` so `TestScan` exercises the new resource type. PRs adding a
mapping without a fixture will be asked to add one before merge.

### Pre-commit Hooks

The repository uses [pre-commit](https://pre-commit.com). Install and run the hooks before submitting:

```bash
pre-commit install
pre-commit run --all-files
```

## Vulnerability Reports

Please do not open public issues for security vulnerabilities. See [SECURITY.md](SECURITY.md) for the private reporting process.

## Licence

By contributing you agree that your changes will be licensed under the Apache License 2.0.
