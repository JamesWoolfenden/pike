# Azure Permissions Data Collection

This directory contains scripts to collect comprehensive Azure RBAC permissions data for generating Pike datasource mappings.

## Overview

Pike needs Azure permissions data to generate accurate IAM mappings for Terraform datasources. This automation collects all available Azure resource provider operations using Azure PowerShell and converts them into a structured format.

## Prerequisites

- **Azure PowerShell Module**: Az.Resources
- **Python 3**: For parsing and conversion

### Installing Azure PowerShell Module

```powershell
Install-Module -Name Az.Resources -Scope CurrentUser
```

Or if you have the full Az module:

```powershell
Install-Module -Name Az -Scope CurrentUser
```

## Usage

### Step 1: Export Azure Permissions

Run the PowerShell script to export all Azure provider operations:

```powershell
cd terraform/azurerm
./export_azure_permissions.ps1
```

This will create `azure_permissions_full.json` containing all Azure resource provider operations.

**Note**: You do NOT need to be logged into Azure. The `Get-AzProviderOperation` cmdlet returns metadata about available operations without requiring authentication.

### Step 2: Parse and Structure the Data

Run the Python script to convert the raw export into Pike's format:

```bash
python3 parse_azure_permissions.py
```

This will create `azure_permissions.json` with permissions organized by:

- Provider (e.g., Microsoft.Storage)
- Resource type (e.g., storageAccounts)
- Permission category (read, write, delete, action)

### Step 3: Use with the Generator

The `azure_permissions.json` file can now be used by `azure_mapping_generator.py` to generate datasource mappings with complete permission coverage.

## File Structure

```shell
terraform/azurerm/
├── export_azure_permissions.ps1  # PowerShell export script
├── parse_azure_permissions.py    # Python parsing script
├── azure_permissions_full.json   # Raw PowerShell output (gitignored)
├── azure_permissions.json        # Structured permissions data (gitignored)
└── azure_mapping_generator.py    # Main generator script
```

## Output Format

The `azure_permissions.json` file is structured as:

```json
{
  "Microsoft.Storage": {
    "storageAccounts": {
      "read": [
        "Microsoft.Storage/storageAccounts/read"
      ],
      "write": [
        "Microsoft.Storage/storageAccounts/write"
      ],
      "delete": [
        "Microsoft.Storage/storageAccounts/delete"
      ],
      "action": [
        "Microsoft.Storage/storageAccounts/listKeys/action",
        "Microsoft.Storage/storageAccounts/regenerateKey/action"
      ]
    }
  }
}
```

## Expected Results

- **Providers**: 100+ Azure resource providers
- **Operations**: 10,000+ individual permissions
- **Coverage**: Complete Azure RBAC operation set

This is a significant improvement over the previous ~27 providers with limited coverage.

## Datasource Generation

For datasources (read-only operations), the generator will use:

- ✅ Permissions from `read` category
- ✅ Some permissions from `action` category (list*, get*)
- ❌ Never permissions from `write` or `delete` categories

## Troubleshooting

### PowerShell Module Not Found

```pwsh
Install-Module -Name Az.Resources -Scope CurrentUser
```

### Python JSON Parsing Errors

Ensure the PowerShell script completed successfully and created `azure_permissions_full.json`.

### Permission Count Seems Low

The `Get-AzProviderOperation` cmdlet returns all publicly documented Azure operations. If the count seems low, verify you're using the latest version of the Az.Resources module:

```powershell
Update-Module -Name Az.Resources
```

## Maintenance

Azure adds new resource providers and operations regularly. To refresh the permissions data:

1. Run `export_azure_permissions.ps1` to get the latest operations
2. Run `parse_azure_permissions.py` to regenerate the structured file
3. Test datasource generation with the updated data

## References

- [Azure RBAC Provider Operations](https://learn.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations)
- [Get-AzProviderOperation Documentation](https://learn.microsoft.com/en-us/powershell/module/az.resources/get-azprovideroperation)
- [Azure Resource Providers](https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/resource-providers-and-types)
