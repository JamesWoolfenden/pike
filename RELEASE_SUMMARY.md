# Pike Release Script - Summary

## What Was Done

### 1. Fixed `parse_azure_permissions.py`

- Added better error handling for empty/invalid JSON files
- Now provides clear instructions when the input file is missing or empty

### 2. Fixed `export_azure_permissions.ps1`

- Added validation to check if Azure operations were retrieved
- Fails early with helpful message if not logged into Azure
- Prevents creation of empty JSON files

### 3. Enhanced `resource.ps1` (Multi-Provider)

- **Now supports all providers** (AWS, Azure, GCP) - auto-detects from resource name
- Fixed PowerShell JSON parsing error using `-AsHashtable` flag
- Updated to handle duplicate keys in Azure permissions data (Microsoft.AWSConnector vs Microsoft.AwsConnector)
- Added provider validation with helpful error messages
- Azure permission lookup only runs for Azure resources
- Updated next steps to reference `release.ps1` with provider parameter

### 4. Created `release.ps1` (Multi-Provider)

**NEW** comprehensive script that automatically updates Go embed files for **all cloud providers**.

**Supports:**

- ✅ AWS (1469 resources, 578 datasources)
- ✅ Azure/azurerm (49 resources, 109 datasources)
- ✅ GCP/google (749 resources, 339 datasources)

**Features:**

- Scans all JSON mapping files in `src/mapping/{provider}/{resource|data}/`
- Generates embed files with `//go:embed` directives
- Updates lookup maps in provider Go files
- Shows detailed statistics (new/existing/removed entries)
- Supports `-DryRun` mode to preview changes
- Supports `-Verbose` for detailed output
- Supports `-Provider` to update just one provider

**Usage:**

```powershell
# Preview all providers
pwsh release.ps1 -DryRun -Verbose

# Update all providers
pwsh release.ps1

# Update only Azure
pwsh release.ps1 -Provider azurerm

# Update only AWS
pwsh release.ps1 -Provider aws

# Update only GCP
pwsh release.ps1 -Provider google
```

**Files Updated:**

| Provider  | Files Updated                                                                                                        |
|-----------|----------------------------------------------------------------------------------------------------------------------|
| **AWS**   | `files_aws.go`, `files_datasource.go`, `aws.go` (tFLookup), `aws_datasource.go` (tFLookupDataAWS)                    |
| **Azure** | `files_azure.go`, `files_azure_datasource.go`, `azure.go` (tFLookupAzure), `azure_datasource.go` (TFLookupAzureData) |
| **GCP**   | `files_gcp.go`, `files_gcp_datasource.go`, `gcp.go` (gCPTfLookup), `gcp_datasource.go` (TFLookup)                    |

### 5. Updated Documentation

- Created `terraform/azurerm/RESOURCE_SCRIPT_README.md` with complete workflow documentation
- Documented all scripts and their usage
- Included troubleshooting section

## Workflow Summary

### For Any Provider (AWS, Azure, GCP)

1. **Optional - Azure permissions setup (Azure only, one-time):**

   ```bash
   cd terraform/azurerm
   Connect-AzAccount
   pwsh export_azure_permissions.ps1
   python3 parse_azure_permissions.py
   ```

2. **Add new resource (any provider):**

   ```bash
   cd /path/to/pike

   # AWS
   pwsh resource.ps1 -resource aws_s3_bucket

   # Azure (with automatic permission lookup)
   pwsh resource.ps1 -resource azurerm_storage_account

   # GCP
   pwsh resource.ps1 -resource google_storage_bucket

   # Edit the generated JSON mapping file
   pwsh release.ps1 -Provider <provider>
   go build && go test ./...
   ```

### For All Providers

```bash
# After adding/modifying any JSON mappings
pwsh release.ps1
go build && go test ./...
git add . && git commit -m "Update provider mappings"
```

## Key Benefits

1. **No more manual editing** of Go embed files
2. **No more manual updating** of lookup maps
3. **Multi-provider support** - one script for all clouds
4. **Safety with dry-run** mode
5. **Clear statistics** on what changed
6. **Automatic permission lookup** for Azure resources

## Files Modified

- ✅ `parse_azure_permissions.py` - Better error handling
- ✅ `export_azure_permissions.ps1` - Validation added
- ✅ `resource.ps1` - **Enhanced to support all providers** (AWS, Azure, GCP)
- ✅ `release.ps1` - **NEW** multi-provider automation script
- ✅ `terraform/azurerm/RESOURCE_SCRIPT_README.md` - **NEW** comprehensive docs
- ✅ `RELEASE_SUMMARY.md` - **NEW** this summary document

## Next Steps

1. Test the `release.ps1` script on your actual mappings
2. Run `pwsh release.ps1 -DryRun` to see current state
3. Optionally run `pwsh release.ps1` to update all Go files
4. Run `go build` to verify everything compiles
5. Commit the changes when ready
