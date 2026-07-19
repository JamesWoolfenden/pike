# Description

## Checklist

- [ ] Tests pass (`make test`)
- [ ] Lint clean (`golangci-lint run`)
- [ ] gosec clean (`make gosec`)
- [ ] New resource mappings include a matching
      `terraform/<provider>/backup/<type>.tf` fixture
- [ ] Golden files regenerated if CLI/API surface changed
      (`make goldens`)
