# Azure Resource Mapping Workflow

This directory contains scripts to help manage Azure resource and datasource mappings for Pike.

## Scripts Overview

**Note:** Scripts 1-2 are **Azure-specific**, while `resource.ps1` and `release.ps1` support **all providers** (AWS, Azure, GCP).

### 1. `export_azure_permissions.ps1` (Azure only)

Exports all Azure resource provider operations from Azure to a JSON file.

**Prerequisites:**

- Az.Resources PowerShell module installed
- Logged into Azure with `Connect-AzAccount`

**Usage:**

```powershell
pwsh export_azure_permissions.ps1
```

**Output:** `azure_permissions_full.json`

### 2. `parse_azure_permissions.py` (Azure only)

Parses the exported Azure permissions and organizes them by provider and resource type.

**Prerequisites:**

- Python 3
- `azure_permissions_full.json` must exist (run script #1 first)

**Usage:**

```bash
python3 parse_azure_permissions.py
```

**Output:** `azure_permissions.json`

### 3. `resource.ps1` (in pike root) - Multi-Provider

Creates a new resource or datasource mapping for **any provider** (AWS, Azure, GCP).

**Features:**

- **Multi-provider support** - automatically detects provider from resource name
- Validates resource names against provider members
- Looks up Azure permissions automatically (Azure only)
- Creates JSON mapping from template
- Creates Terraform test file
- Provides interactive permission selection for Azure resources

**Usage:**

```powershell
# AWS resources
pwsh ../../resource.ps1 -resource aws_s3_bucket
pwsh ../../resource.ps1 -resource aws_lambda_function -type data

# Azure resources (with automatic permission lookup)
pwsh ../../resource.ps1 -resource azurerm_storage_account
pwsh ../../resource.ps1 -resource azurerm_storage_account -type data

# GCP resources
pwsh ../../resource.ps1 -resource google_storage_bucket
pwsh ../../resource.ps1 -resource google_compute_instance -type data
```

**What it does:**

1. Auto-detects provider from resource name prefix (aws_, azurerm_, google_)
2. Validates the resource name exists in the provider
3. Looks up Azure permissions from `azure_permissions.json` (Azure only)
4. Shows permission counts (read/write/delete/action) for Azure
5. Copies template.json to the new resource JSON file
6. Creates a Terraform test file

### 4. `release.ps1` (in pike root)

Automatically updates Go embed files and lookup maps for **all cloud providers** (AWS, Azure, GCP).

**Features:**

- Scans all resource and datasource JSON mappings for all providers
- Generates embed files (`files_*.go`)
- Updates lookup maps in provider-specific Go files
- Shows what's new, changed, or removed
- Supports processing all providers or a single provider

**Usage:**

```powershell
# Dry run to see what would change for ALL providers
pwsh ../../release.ps1 -DryRun

# Verbose output showing specific changes
pwsh ../../release.ps1 -DryRun -Verbose

# Update only Azure files
pwsh ../../release.ps1 -Provider azurerm

# Update only AWS files
pwsh ../../release.ps1 -Provider aws

# Update only GCP files
pwsh ../../release.ps1 -Provider google

# Actually update all providers
pwsh ../../release.ps1
```

**What it updates:**

**AWS:**

- `src/files_aws.go` (resource embeds)
- `src/files_datasource.go` (datasource embeds)
- `src/aws.go` (tFLookup map)
- `src/aws_datasource.go` (tFLookupDataAWS map)

**Azure (azurerm):**

- `src/files_azure.go` (resource embeds)
- `src/files_azure_datasource.go` (datasource embeds)
- `src/azure.go` (tFLookupAzure map)
- `src/azure_datasource.go` (TFLookupAzureData map)

**GCP (google):**

- `src/files_gcp.go` (resource embeds)
- `src/files_gcp_datasource.go` (datasource embeds)
- `src/gcp.go` (gCPTfLookup map)
- `src/gcp_datasource.go` (TFLookup map)

## Complete Workflow

### One-Time Setup

1. **Export Azure permissions:**

   ```powershell
   cd terraform/azurerm
   Connect-AzAccount  # Login to Azure
   pwsh export_azure_permissions.ps1
   ```

2. **Parse the permissions:**

   ```bash
   python3 parse_azure_permissions.py
   ```

This creates `azure_permissions.json` which the resource script will use for automatic permission lookup.

### Adding a New Resource or Datasource

1. **Create the resource mapping:**

   ```powershell
   cd /path/to/pike
   pwsh resource.ps1 -resource azurerm_storage_account
   # or for datasource:
   pwsh resource.ps1 -resource azurerm_storage_account -type data
   ```

2. **Edit the JSON mapping:**

   - Open the created JSON file (shown in output)
   - Add the Azure permissions that were displayed
   - Configure the attribute mappings as needed

3. **Update the Go files:**

   ```powershell
   pwsh release.ps1
   ```

4. **Build and test:**

   ```bash
   go build
   go test ./...
   ```

5. **Commit your changes:**

   ```bash
   git add .
   git commit -m "Add support for azurerm_storage_account"
   ```

## Directory Structure

```text
pike/
├── resource.ps1                              # Create new resource mappings
├── release.ps1                               # Update Go embed files
├── src/
│   ├── mapping/azurerm/
│   │   ├── resource/                         # Resource mappings
│   │   │   ├── storage/
│   │   │   │   └── azurerm_storage_account.json
│   │   │   └── template.json
│   │   └── data/                             # Datasource mappings
│   │       ├── storage/
│   │       │   └── azurerm_storage_account.json
│   │       └── template.json
│   ├── files_azure.go                        # Resource embeds (auto-generated)
│   ├── files_azure_datasource.go             # Datasource embeds (auto-generated)
│   ├── azure.go                              # Resource lookup map (auto-updated)
│   └── azure_datasource.go                   # Datasource lookup map (auto-updated)
└── terraform/azurerm/
    ├── export_azure_permissions.ps1
    ├── parse_azure_permissions.py
    ├── azure_permissions_full.json           # Raw Azure data
    └── azure_permissions.json                # Parsed permissions
```

## Tips

- **Both scripts support all providers** (AWS, Azure, GCP) - provider is auto-detected from resource name
- Run `release.ps1 -DryRun -Verbose` to preview changes before applying
- The `resource.ps1` script will suggest similar resource names if you typo
- Azure permissions are automatically looked up for Azure resources (not available for AWS/GCP)
- Use fuzzy matching for Azure permissions - the script will show options if multiple matches are found
- Always run `release.ps1` after adding new JSON mappings to update the Go code
- Use `release.ps1 -Provider <name>` to update just one provider if you only changed one
- The release script processes **1469 AWS resources, 49 Azure resources, and 749 GCP resources**

## Troubleshooting

### "azure_permissions.json not found"

- Run the export and parse scripts first (steps 1-2 in One-Time Setup)

### "Resource not found in provider"

- Check the resource name spelling
- Verify the resource exists in Terraform's azurerm provider
- Look at suggestions shown by the script

### "No Azure permissions found"

- The resource name might not match any Azure resource type
- Try a different resource name or check the permissions file manually

### PowerShell errors with JSON parsing

- This is fixed in the updated `resource.ps1` with `-AsHashtable` flag
- Make sure you're using the latest version of the script
