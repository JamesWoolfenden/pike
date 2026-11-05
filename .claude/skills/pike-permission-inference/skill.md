---
name: pike-permission-inference
description: Infer GCP IAM permissions for a new Pike resource mapping by pattern-matching against similar resources, then validate the permissions exist via gcloud role creation. Use when you can't or don't want to create actual cloud resources. Inferior to the empirical method (pike-mapping-discovery) but faster for obvious cases.
---

# pike-permission-inference

Lightweight alternative to `pike-mapping-discovery` for adding Pike mappings when:
- The resource is well-documented with obvious permission patterns
- You can't provision the actual resource (cost, complexity, dependencies)
- You want a quick first pass that can be refined later

**Trade-off:** This proves permissions *exist* in GCP's IAM system, but not that they're the *correct* ones for Terraform's lifecycle. If wrong, users hit 403s at apply/destroy time.

---

## Process

### 1. Check if the Terraform resource actually exists

Before inferring permissions, verify the resource is real:

```bash
# Search Terraform registry
# Look for direct hits on registry.terraform.io

# Check provider GitHub for recent additions
# Search: github terraform-provider-google <resource_name>
```

If the resource doesn't exist in the provider, delete the `.tf` file referencing it rather than creating a mapping.

### 2. Find similar existing mappings

Look for resources in the same GCP service:

```bash
# List existing mappings for the service
ls src/mapping/google/resource/<service>/

# Read a similar resource's mapping
cat src/mapping/google/resource/<service>/google_<similar>.json
```

GCP permission naming follows `<service>.<resourcePlural>.<verb>`:
- `compute.instances.create` / `compute.instances.delete` / etc.
- `monitoring.alertPolicies.create` / `monitoring.alertPolicies.update` / etc.

### 3. Infer permissions from naming patterns

Standard verbs for resources:
- `plan`: `.get` (read current state)
- `apply`: `.create`, `.get`, `.update`, `.delete` (full lifecycle)
- `modify`: `.update`
- `destroy`: `.delete`

Data sources typically only need:
- `plan`/`apply`: `.get` or `.list`

**Watch for exceptions:**
- Some resources use `.cancel` instead of `.delete` (e.g., `monitoring.snoozes`)
- Regional resources often have different permission names (e.g., `compute.regionSslPolicies.*` vs `compute.sslPolicies.*`)

### 4. Verify permissions exist via web search

Search for the exact permission names:
- `site:gcp.permissions.cloud "<permission.name>"`
- `site:cloud.google.com/iam/docs "<permission.name>"`
- Check which predefined roles include the permission

### 5. Validate permissions against GCP

Create a test role with the inferred permissions:

```bash
gcloud iam roles create pike_test_validate --project=pike-477416 \
  --permissions="<perm1>,<perm2>,<perm3>" \
  --title="Pike Test Validate"
```

**If GCP accepts the role:** permissions are valid (exist in GCP's IAM system).

**If GCP rejects with "invalid permission":** one or more permissions don't exist - research further.

Clean up after:
```bash
gcloud iam roles delete pike_test_validate --project=pike-477416 --quiet
```

### 6. Write the mapping file

Create the JSON mapping in the appropriate directory:

```bash
# Determine directory from permission prefix
# compute.* → src/mapping/google/resource/compute/
# monitoring.* → src/mapping/google/resource/monitoring/
```

Standard resource mapping structure:
```json
[
  {
    "apply": [
      "<service>.<resource>.create",
      "<service>.<resource>.delete",
      "<service>.<resource>.get",
      "<service>.<resource>.update"
    ],
    "attributes": {
      "tags": []
    },
    "destroy": [
      "<service>.<resource>.delete"
    ],
    "modify": [
      "<service>.<resource>.update"
    ],
    "plan": [
      "<service>.<resource>.get"
    ]
  }
]
```

### 7. Rebuild and verify

```bash
make build
./pike scan -d terraform/google  # verify permissions are picked up
```

---

## When NOT to use this method

- When the resource has complex attribute-dependent permissions (use `pike-mapping-discovery`)
- When you need 100% confidence in the permissions
- When the permission pattern isn't obvious from similar resources
- When web searches don't confirm the permission names exist

In these cases, fall back to `pike-mapping-discovery` which empirically tests against real GCP API responses.

---

## Example session

Adding `google_compute_region_ssl_policy`:

1. **Similar resource exists:** `google_compute_ssl_policy` uses `compute.sslPolicies.*`
2. **Data source already exists:** `data.google_compute_region_ssl_policy` uses `compute.regionSslPolicies.get`
3. **Infer resource permissions:** `compute.regionSslPolicies.{create,delete,get,update}`
4. **Web search confirms:** Found in Compute Network Admin role on gcp.permissions.cloud
5. **GCP accepts test role:** `gcloud iam roles create` succeeded
6. **Write mapping:** `src/mapping/google/resource/compute/google_compute_region_ssl_policy.json`
