# Azure Datasource Generation Requirements

## Critical Rules

### 1. Datasources are Read-Only

Datasources MUST have:

```json
{
  "apply": [],    // Always empty - datasources never apply changes
  "destroy": [],  // Always empty - datasources never destroy resources
  "modify": []    // Always empty - datasources never modify resources
}
```

### 2. Plan Permissions

Datasources can have in `plan` array:

- ✅ `/read` permissions - Standard read operations
- ✅ `/action` permissions - Read-like operations (e.g., listKeys, listSecrets)
- ❌ `/write` permissions - NEVER for datasources
- ❌ `/delete` permissions - NEVER for datasources

**Example:**

```json
"plan": [
  "Microsoft.Storage/storageAccounts/read",           // ✅ read
  "Microsoft.Storage/storageAccounts/listKeys/action"  // ✅ action (read-like)
]
```

### 3. Permission Types for Datasources

**ALLOWED:**

- `*/read` - Read resource properties
- `*/list*` - List operations (subresources, keys, secrets, etc.)
- `*/get*` - Get specific items
- Some `*/action` - Read-like actions (listKeys, getStatus, etc.)

**NOT ALLOWED:**

- `*/write` - Create or update
- `*/delete` - Delete resources
- `*/action` for destructive operations

### 4. Hierarchical Permissions

Datasources may need parent resource permissions:

```json
"plan": [
  "Microsoft.Resources/subscriptions/resourcegroups/read",  // Parent
  "Microsoft.Network/virtualNetworks/read",                 // Parent
  "Microsoft.Network/virtualNetworks/subnets/read"          // Target resource
]
```

## Generator Implementation Requirements

### Method Signature

```python
def generate_mapping(self, tf_resource: str, is_datasource: bool = False) -> Dict:
```

### Logic for Datasources

```python
if is_datasource:
    # Filter to read-only permissions
    read_like_perms = [p for p in resource_perms
                       if '/read' in p.lower() or
                          '/list' in p.lower() or
                          '/get' in p.lower()]

    return [{
        "apply": [],
        "attributes": {"tags": []},
        "destroy": [],
        "modify": [],
        "plan": sorted(read_like_perms)
    }]
else:
    # Resources get full permissions (write, delete, etc.)
    return [{
        "apply": sorted(write_perms + read_perms + delete_perms + action_perms),
        "destroy": sorted(delete_perms),
        "modify": [],
        "plan": sorted(read_perms)
    }]
```

## Testing Strategy

1. **Keep baseline samples** in `test/datasource_samples/`
2. **Test generator** against known-good datasources
3. **Verify output structure**:
   - Empty apply/destroy/modify
   - Only read-like permissions in plan
   - No write/delete permissions

## Known Limitations

1. **Azure permissions file incomplete**
   - Only ~50 providers have permission data
   - Missing: Microsoft.Dns, Microsoft.Virtual, Microsoft.Vpn, etc.
   - Cannot generate mappings without Azure permission data

2. **Realistic expectations**
   - Can only generate ~50-60 datasources with current data
   - Need to expand `azure_permissions.json` for more coverage

3. **Manual curation needed**
   - Some `/action` permissions may be ambiguous
   - May need manual review for edge cases
   - Consider allowlist/denylist for action permissions

## Next Steps

1. ✅ Preserve test samples
2. ⏹️ Fix generator to support `is_datasource` parameter
3. ⏹️ Test against baseline samples
4. ⏹️ Generate only datasources with confident permission data
5. ⏹️ Skip datasources without Azure provider data
6. ⏹️ Verify compilation
7. ⏹️ Commit working state
