# Datasource Test Samples

These are baseline datasource mappings used to verify the generator produces correct output.

## Test Cases

### 1. azurerm_storage_account.json

**Provider:** Microsoft.Storage
**Complexity:** Simple (2 permissions)
**Permissions:**

- `Microsoft.Storage/storageAccounts/read`
- `Microsoft.Storage/storageAccounts/listKeys/action`

**Notes:**

- Includes an `/action` permission (listKeys) - datasources can have action permissions for read operations
- Good test for simple storage resources

### 2. azurerm_key_vault.json

**Provider:** Microsoft.KeyVault
**Complexity:** Simple (2 permissions)
**Permissions:**

- `Microsoft.Resources/subscriptions/resourcegroups/read`
- `Microsoft.KeyVault/vaults/read`

**Notes:**

- Includes parent resource group read permission
- Good test for hierarchical permissions

### 3. azurerm_subnet.json

**Provider:** Microsoft.Network
**Complexity:** Medium (3 permissions)
**Permissions:**

- `Microsoft.Resources/subscriptions/resourcegroups/read`
- `Microsoft.Network/virtualNetworks/read`
- `Microsoft.Network/virtualNetworks/subnets/read`

**Notes:**

- Shows hierarchical permission structure (resource group → vnet → subnet)
- Good test for nested resources

### 4. azurerm_api_management.json

**Provider:** Microsoft.ApiManagement
**Complexity:** Simple (1 permission)
**Permissions:**

- `Microsoft.ApiManagement/service/read`

**Notes:**

- Simplest case - single read permission
- Good baseline test

## Expected Structure

All datasources must have:

```json
[
  {
    "apply": [],      // MUST be empty - datasources don't apply changes
    "attributes": {
      "tags": []
    },
    "destroy": [],    // MUST be empty - datasources don't destroy
    "modify": [],     // MUST be empty - datasources don't modify
    "plan": [...]     // Contains read and action permissions
  }
]
```

## Key Requirements

1. ✅ **Read-only**: `apply`, `destroy`, `modify` must be empty arrays
2. ✅ **Plan permissions**: Can include `/read` and `/action` permissions
3. ✅ **No write/delete**: Should NOT include `/write` or `/delete` permissions
4. ✅ **Hierarchical**: Can include parent resource permissions (resource groups, parent resources)

## Testing the Generator

To test if the generator works correctly:

```bash
# Generate a test datasource
python3 terraform/azurerm/azure_mapping_generator.py --mode generate --type datasource --resources azurerm_storage_account

# Compare output
diff src/mapping/azurerm/data/storage/azurerm_storage_account.json test/datasource_samples/azurerm_storage_account.json
```

The generator should produce identical or semantically equivalent output (same permissions, may differ in order).
