# Contributing to Pike

Thank you for your interest in contributing to Pike! Contributions are welcome and appreciated.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Environment Setup](#development-environment-setup)
- [Adding New Resource Mappings](#adding-new-resource-mappings)
- [Testing Your Changes](#testing-your-changes)
- [Code Style Guidelines](#code-style-guidelines)
- [Submitting Pull Requests](#submitting-pull-requests)
- [Getting Help](#getting-help)

## Getting Started

Before contributing, please:

1. Check existing [issues](https://github.com/JamesWoolfenden/pike/issues) to see if your bug/feature is already being discussed
2. For new features, open an issue first to discuss the approach
3. Fork the repository and create a branch for your work

## Development Environment Setup

### Prerequisites

- **Go 1.19+** - [Install Go](https://golang.org/doc/install)
- **Git** - For version control
- **pre-commit** - For running automated checks (required)
- **PowerShell (pwsh)** - For resource generation scripts: `brew install --cask powershell`

### Initial Setup

1. **Clone the repository**

   ```shell
   git clone https://github.com/JamesWoolfenden/pike.git
   cd pike
   ```

2. **Install dependencies**

   ```shell
   go mod download
   ```

3. **Build Pike**

   ```shell
   go build -o pike
   ```

   Or using Make:

   ```shell
   make build
   ```

4. **Install pre-commit framework** (required)

   ```shell
   # macOS/Linux
   brew install pre-commit

   # Or using pip
   pip install pre-commit
   ```

5. **Set up pre-commit hooks**

   ```shell
   pre-commit install
   ```

   All code submissions must pass pre-commit checks before being merged.

### Running Pike Locally

Test your changes:

```shell
# Scan a local directory
./pike scan -d ./terraform/aws

# Run with different output formats
./pike scan -o terraform -d ./terraform/aws
```

## Adding New Resource Mappings

Pike supports AWS, Azure, and GCP. Adding support for new resources involves creating IAM permission mappings.

### Overview

Each cloud resource needs a mapping file that defines the minimum IAM permissions required for:

- **apply** - Creating the resource
- **destroy** - Deleting the resource
- **plan** - Reading/describing the resource
- **modify** - Updating the resource
- **attributes** - Additional permissions based on resource attributes

### Mapping File Structure

Mapping files are JSON files located in `./src/mapping/<provider>/<type>/`.

Example: `aws_security_group.json`:

```json
[
  {
    "apply": [
      "ec2:CreateSecurityGroup",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeAccountAttributes",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DeleteSecurityGroup",
      "ec2:RevokeSecurityGroupEgress"
    ],
    "attributes": {
      "ingress": [
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:AuthorizeSecurityGroupEgress"
      ],
      "tags": [
        "ec2:CreateTags",
        "ec2:DeleteTags"
      ]
    },
    "destroy": [
      "ec2:DeleteSecurityGroup"
    ],
    "modify": [],
    "plan": []
  }
]
```

### Step-by-Step: Adding a New Resource

#### 1. Generate the Mapping Template

Use the resource.ps1 script to create boilerplate files:

```shell
./resource.ps1 aws_new_resource_name
```

This creates:

- A blank mapping file in `src/mapping/`
- A test .tf file in `terraform/<provider>/`

#### 2. Determine Required Permissions

**Datasources are the easiest to start with.**

Create a minimal resource/datasource .tf file and place it in the correct directory (e.g., `terraform/aws/`).

Run OpenTofu/Terraform with minimal permissions to trigger permission errors:

```shell
cd terraform/aws
tofu init
tofu plan
```

The error output will indicate which permissions are missing. Common sources:

- OpenTofu/Terraform debug output
- AWS IAM Policy Simulator
- Azure Activity Logs
- GCP Cloud Audit Logs

**Note:** Not all providers clearly indicate required permissions. You may need to:

- Check official IAM documentation
- Use cloud provider policy makers/simulators
- Trial and error with permission additions

**Common gotchas:**

- CloudFormation-based resources may require additional permissions not explicitly stated
- Some resources require `iam:PassRole` or `iam:CreateServiceLinkedRole` without indicating it
- Wildcards may be needed initially, then refined

#### 3. Understand Attribute-Based Permissions

Some resources require extra permissions depending on attributes used:

```json
"attributes": {
  "ingress": ["ec2:AuthorizeSecurityGroupIngress"],
  "tags": ["ec2:CreateTags", "ec2:DeleteTags"],
  "encryption": ["kms:Decrypt", "kms:Encrypt"]
}
```

Build comprehensive .tf examples covering all common attributes.

#### 4. Handle Eventual Consistency

**AWS:** IAM roles can take time to propagate (typically seconds to minutes)

**Azure:** Resource operations can be slow; allow extra time for testing

**GCP:** Generally faster, but some operations may still have delays

#### 5. Add the Import to files.go

Update `src/files.go` to embed your mapping file:

```go
//go:embed aws_security_group.json
var securityGroup []byte
```

#### 6. Update the Provider Lookup Table

Add your resource to the appropriate provider scan function.

For AWS, update `GetAWSResourcePermissions()`:

```go
func GetAWSResourcePermissions(result template) []interface{} {
    TFLookup := map[string]interface{}{
        "aws_s3_bucket":            awsS3Bucket,
        "aws_s3_bucket_acl":        awsS3BucketACL,
        "aws_security_group":       awsSecurityGroup,  // Add this line
        // ...
    }
```

For Azure: Update `GetAzureRMResourcePermissions()`
For GCP: Update `GetGoogleResourcePermissions()`

#### 7. Add Test Cases

Add an example .tf file to `terraform/<cloud>/backups/` to ensure your mapping is picked up by pike:

```hcl
# terraform/aws/backups/aws_new_resource.tf
resource "aws_new_resource" "example" {
  name = "test"
  # Add realistic configuration
}
```

## Testing Your Changes

### Manual Testing

```shell
# Build pike
go build -o pike

# Test against your new resource
./pike scan -d terraform/aws/backups

# Verify the output includes your new permissions
./pike scan -o terraform -d terraform/aws/backups
```

### Automated Tests

Run the test suite:

```shell
go test ./...
```

Run pre-commit checks:

```shell
pre-commit run --all-files
```

### Coverage Files

Pike tracks resource coverage in `src/coverage/*.md`. These files are **auto-generated** and show which resources still need mappings. After adding a new resource, the coverage percentage should increase.

## Code Style Guidelines

- **Go code**: Follow standard Go conventions (`gofmt`, `golint`)
- **JSON files**: Use 2-space indentation, alphabetize permissions within each section
- **Comments**: Add comments for non-obvious permission requirements
- **Commit messages**: Use clear, descriptive messages
  - Format: `<type>: <description>`
  - Types: `feat`, `fix`, `docs`, `test`, `refactor`, `chore`
  - Example: `feat: add support for aws_lambda_function`

## Submitting Pull Requests

1. **Ensure your code passes all checks**

   ```shell
   pre-commit run --all-files
   go test ./...
   ```

2. **Create a pull request** with:
   - Clear title describing the change
   - Description explaining what was added/changed
   - Reference any related issues (e.g., "Fixes #123")
   - Screenshots or example output if applicable

3. **PR Checklist**:
   - [ ] Pre-commit hooks pass
   - [ ] Tests pass
   - [ ] Mapping file(s) added to `src/mapping/`
   - [ ] files.go updated with embed directive
   - [ ] Provider lookup table updated
   - [ ] Test .tf file added to `terraform/<cloud>/backups/`
   - [ ] Manually tested with `pike scan`

4. **Review process**:
   - Maintainers will review your PR
   - Address any feedback or requested changes
   - Once approved, your PR will be merged

## Getting Help

- **Questions?** Open a [discussion](https://github.com/JamesWoolfenden/pike/discussions)
- **Bug reports?** Open an [issue](https://github.com/JamesWoolfenden/pike/issues)
- **Contact:** James Woolfenden <james.woolfenden@gmail.com>

---

Thank you for contributing to Pike! Your efforts help make infrastructure-as-code more secure for everyone.
