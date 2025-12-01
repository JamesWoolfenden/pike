# resource.ps1 - Enhanced Resource Scaffolding Script

The `resource.ps1` script helps create skeleton files for new Terraform resource/datasource mappings in Pike.

## Features

### 1. Resource Name Validation

Automatically validates that the resource exists in the Terraform provider schema before creating files.

- Checks against `src/parse/{provider}-members.json`
- Shows helpful suggestions if resource name is invalid
- Prevents creating mappings for non-existent resources

### 2. Interactive Permission Selection (Azure only)

When `azure_permissions.json` exists, the script will:

- Search for matching Azure resource types
- Present multiple options if found
- Show permission counts for each option (read/write/delete/action)
- Let you choose which Azure resource type to use
- Display a summary of available permissions

### 3. File Creation

Creates the necessary files for Pike:

- JSON mapping file (from template)
- Terraform test file (.tf)

## Prerequisites

- **PowerShell Core** (pwsh)
- **Provider members file**: `src/parse/{provider}-members.json` (for validation)
- **Azure permissions file** (optional): `terraform/azurerm/azure_permissions.json` (for permission lookup)

## Usage

### Basic Usage

```powershell
# Create a resource mapping
./resource.ps1 -resource azurerm_storage_account

# Create a datasource mapping
./resource.ps1 -resource azurerm_storage_account -type data
```

### Example: Creating an Azure Datasource

```powershell
./resource.ps1 -resource azurerm_key_vault -type data
```

**Output:**

```terminaloutput
Validating resource...
✓ Resource 'azurerm_key_vault' validated in azurerm provider

Looking up permissions...
Found permissions for: Microsoft.KeyVault/vaults

Permission Summary:
  Read permissions:   1
  Write permissions:  1
  Delete permissions: 1
  Action permissions: 0

Creating mapping files...
✓ Copied template to: src/mapping/azurerm/data/azurerm_key_vault.json
✓ Created Terraform file: terraform/azurerm/data.azurerm_key_vault.tf

✓ Done!

Next steps:
  1. Edit the JSON mapping to add the permissions shown above
  2. Update Go embed files (src/files_azure_datasource.go, src/azure_datasource.go)
  3. Test with: go build
```

### Example: Invalid Resource Name

```powershell
./resource.ps1 -resource azurerm_fake_thing -type data
```

**Output:**

```output
Validating resource...
✗ ERROR: Resource 'azurerm_fake_thing' not found in azurerm provider

Did you mean one of these?
  - azurerm_arc_machine
  - azurerm_application_gateway
  - azurerm_kubernetes_cluster

Aborting: Resource validation failed
```

### Example: Multiple Azure Resource Type Matches

```powershell
./resource.ps1 -resource azurerm_storage_account -type data
```

**Output:**

```terminaloutput
Validating resource...
✓ Resource 'azurerm_storage_account' validated in azurerm provider

Looking up permissions...

Found multiple possible Azure resource types:
  [1] Microsoft.Storage/storageAccounts
      Read: 1 | Write: 1 | Delete: 1 | Action: 0
  [2] Microsoft.Storage/storageAccounts/blobServices/containers
      Read: 1 | Write: 0 | Delete: 0 | Action: 0
  [3] Microsoft.Storage/storageAccounts/listKeys
      Read: 0 | Write: 0 | Delete: 0 | Action: 1

Which resource type should we use? [1-3, or 0 to skip]: 1
✓ Selected: Microsoft.Storage/storageAccounts

Permission Summary:
  Read permissions:   1
  Write permissions:  1
  Delete permissions: 1
  Action permissions: 0
```

## Setting Up Azure Permissions

To enable the interactive permission selection feature:

```powershell
# 1. Export Azure permissions
cd terraform/azurerm
./export_azure_permissions.ps1

# 2. Parse and structure the data
python3 parse_azure_permissions.py

# 3. Now resource.ps1 will use azure_permissions.json for permission lookup
```

See `terraform/azurerm/AZURE_PERMISSIONS_README.md` for details.

## Files Created

### JSON Mapping File

Location: `src/mapping/{provider}/{type}/{resource}.json`

Copied from the template with empty permission arrays. You need to manually add permissions.

**For datasources**, the template has the correct structure:

```json
[
  {
    "apply": [],
    "attributes": { "tags": [] },
    "destroy": [],
    "modify": [],
    "plan": []
  }
]
```

### Terraform Test File

**For datasources:** `terraform/{provider}/data.{resource}.tf`

```hcl
data "azurerm_storage_account" "pike" {
}

output "azurerm_storage_account" {
  value = data.azurerm_storage_account.pike
}
```

**For resources:** `terraform/{provider}/{resource}.tf`

```hcl
resource "azurerm_storage_account" "pike" {}
```

## Next Steps After Running the Script

1. **Edit the JSON mapping** to add permissions:
   - For datasources: Add read permissions to the `plan` array
   - For resources: Add write/delete to `apply`, delete to `destroy`, read to `plan`

2. **Update Go embed files** (for datasources):
   - Add `//go:embed` directive in `src/files_azure_datasource.go`
   - Add case entry in `src/azure_datasource.go`

3. **Test compilation:**

```bash
   go build
   ```

4. **Test the mapping:**

```bash
   cd terraform/azurerm
   terraform init
   pike scan -d . -o json
```

## Comparison: Manual vs Automated

| Feature | resource.ps1 | azure_mapping_generator.py |
|---------|--------------|---------------------------|
| Validation | ✅ Yes | ✅ Yes |
| Permission lookup | ✅ Interactive | ✅ Automatic |
| File creation | ✅ One-at-a-time | ✅ Batch mode |
| Go file updates | ❌ Manual | ❌ Manual* |
| Best for | Individual resources | Bulk generation |

\* Future enhancement

## Enhancements Added

- ✅ **Enhancement 3**: Resource name validation against provider schema
- ✅ **Enhancement 5**: Interactive permission selection for Azure resources

## Future Enhancements

Potential improvements not yet implemented:

- Auto-populate JSON with selected permissions
- Automatically update Go embed files
- Support for AWS/Google permission lookup
- Batch mode for multiple resources
