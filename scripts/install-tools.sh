#!/bin/bash
# Install Go development and security tools

set -e

echo "Installing Go development and security tools..."

# Security tools
echo "📦 Installing security scanners..."
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install golang.org/x/vuln/cmd/govulncheck@latest

# Code quality tools
echo "📦 Installing code quality tools..."
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest


# Dependency analysis
echo "📦 Installing dependency analysis tools..."
go install github.com/sonatype-nexus-community/nancy@latest

# OpenAPI code generator (already in go.mod tools)
echo "📦 Installing OpenAPI code generator..."
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

echo ""
echo "✅ All tools installed successfully!"
echo ""
echo "Available make targets:"
echo "  make vet           - Run go vet"
echo "  make staticcheck   - Run staticcheck"
echo "  make gosec         - Run security scanner"
echo "  make govulncheck   - Check vulnerabilities"
echo "  make complexity    - Check code complexity"
echo "  make check-all     - Run all checks"
echo "  make security-scan - Run security-focused checks"
echo ""
echo "Run 'make help' to see all available targets"
