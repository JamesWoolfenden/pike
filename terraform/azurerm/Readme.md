# Azure Tools for Pike

This directory contains tools and scripts for working with Pike and Azure.

## Auth for working on Pike with Azure

Unlike AWS you cant name profiles.

You will need two distinct auth identities, one that has the power to update roles, so that you can update the
permissions of the other role that you use to determine the permissions required to create a resource.
You can create the simple role using the contents of <terraform/azurerm/role>; this is the role that you update.

I work in two shells, each with different permissions.

```shell
# sh
export ARM_CLIENT_ID="00000000-0000-0000-0000-000000000000"
export ARM_CLIENT_SECRET="12345678-0000-0000-0000-000000000000"
export ARM_TENANT_ID="10000000-0000-0000-0000-000000000000"
export ARM_SUBSCRIPTION_ID="20000000-0000-0000-0000-000000000000"
```

<https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/guides/service_principal_client_secret>

---

## Azure Mapping Generator

Scripts for generating Pike mapping files for Azure resources.

### Prerequisites

- **PowerShell** (pwsh) with Azure PowerShell module
- **Python 3.7+**
- **Azure account** with appropriate permissions

### Workflow

#### 1. Export Azure Permissions (PowerShell)

```powershell
# Install Azure PowerShell module
Install-Module -Name Az.Resources -Scope CurrentUser

# Login to Azure
Connect-AzAccount

# Run export script
pwsh export_azure_permissions.ps1
```

Creates `azure_permissions_full.json` with all Azure provider operations.

#### 2. Parse Permissions (Python)

```bash
python parse_azure_permissions.py
```

Creates `azure_permissions.json` with structured permissions organized by provider and resource type.

#### 3. Generate Mapping Files (Python)

```bash
# Preview what will be generated
python azure_mapping_generator.py --dry-run

# Generate mapping files
python azure_mapping_generator.py

# Generate only for specific provider
python azure_mapping_generator.py --filter Microsoft.Compute
```

Creates mapping files in `src/mapping/azurerm/resource/`.

### Scripts

| Script | Purpose |
|--------|---------|
| `export_azure_permissions.ps1` | Export Azure permissions using PowerShell |
| `parse_azure_permissions.py` | Parse raw permissions into structured JSON |
| `azure_mapping_generator.py` | Generate Pike mapping files |
| `validate_datasources.py` | Validate Azure data source mappings |

### Testing Generated Mappings

```bash
# Build Pike
go build -o pike

# Test with Azure Terraform code
./pike scan -d path/to/azure/terraform -p azurerm
```

### Customization

The generator includes mappings for common Azure resources. To add more, edit `RESOURCE_NAME_MAPPINGS` in `azure_mapping_generator.py`:

```python
RESOURCE_NAME_MAPPINGS = {
    "Microsoft.Compute/virtualMachines": "azurerm_virtual_machine",
    "Microsoft.YourService/yourResource": "azurerm_your_resource",
}
```

For detailed documentation, see the docstrings in each script.
