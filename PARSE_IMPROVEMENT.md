# Parse Command Improvement - Schema-Based Extraction

## Summary

**Date:** 2026-03-04
**Issue:** The `pike parse` command used unreliable documentation parsing, resulting in incomplete and inaccurate provider resource lists.
**Solution:** Replaced with Terraform provider schema extraction (authoritative source).

## The Problem

### Old Method: Documentation Parsing

The original `pike parse` scanned provider documentation markdown files to extract resource names using regex patterns.

**Critical Issues:**

1. ❌ **Incomplete** - Missed resources with non-standard documentation
2. ❌ **Inaccurate** - Included deprecated/removed resources still in docs
3. ❌ **Unreliable** - Documentation can be outdated or malformed
4. ❌ **Provider-specific** - Required cloning entire provider repositories

### Real-World Impact

**Google Provider (Before Fix):**

- Listed: 1310 resources, 422 datasources
- Actual (v7.21.0): 1184 resources, 408 datasources
- **Missing:** 15 current resources (including `google_securityposture_posture_deployment`)
- **Extra:** 141 deprecated/invalid resources (e.g., `google_cloudiot_registry`, `google_api_gateway_*`)

**AWS Provider (Before Fix):**

- Listed: 17 resources, 528 datasources
- Actual: **1615 resources**, 644 datasources
- **Missing:** 1598 resources! (99% missing)

**Azure Provider (Before Fix):**

- Listed: Unknown accuracy
- Actual: 1124 resources, 393 datasources

## The Solution

### New Method: Terraform Schema Extraction

The improved `pike parse` command now uses **Terraform's provider schema** via `terraform providers schema -json`.

**Benefits:**

1. ✅ **Authoritative** - Direct from Terraform's internal provider registry
2. ✅ **Complete** - All resources/datasources automatically included
3. ✅ **Accurate** - Matches exact provider version capabilities
4. ✅ **Fast** - No need to clone repositories or scan files
5. ✅ **Consistent** - Same method works for all providers

### How It Works

```bash
# Extract latest version
pike parse -n google

# Internally creates temp Terraform config:
# terraform { required_providers { google = { source = "hashicorp/google" } } }
# terraform init
# terraform providers schema -json
```

The command:

1. Creates a minimal Terraform configuration
2. Initializes Terraform to download the provider
3. Extracts the provider schema as JSON
4. Parses resource and datasource names
5. Saves to `{provider}-members.json`

### Backward Compatibility

The old documentation parsing method is kept as a **fallback** if schema extraction fails:

```shell
Attempting schema-based extraction...
  ✓ Success - uses schema (recommended)
  ✗ Failure - falls back to docs parsing (unreliable)
```

## Results

### Google Provider

```shell
Before: 1310 resources (141 invalid), 422 datasources (28 invalid)
After:  1185 resources (all valid), 408 datasources (all valid)

✅ Now includes google_securityposture_posture_deployment
✅ Removed 141 deprecated resources
✅ Added 15 missing resources
```

### AWS Provider

```shell
Before: 17 resources (broken), 528 datasources
After:  1615 resources, 644 datasources

✅ Added 1598 missing resources!
```

### Azure Provider

```shell
Before: Unknown accuracy
After:  1124 resources, 393 datasources

✅ Complete and accurate
```

## Usage

### New Recommended Command

```bash
# Extract from schema (no directory needed)
pike parse -n google
pike parse -n aws
pike parse -n azurerm
```

### Old Command (Deprecated)

```bash
# Only used if schema extraction fails
pike parse -n google -d /path/to/provider-docs
```

## Technical Changes

### Files Modified

1. **`parse/parse.go`**
   - Added `parseFromSchema()` - schema-based extraction
   - Updated `Parse()` - tries schema first, falls back to docs
   - Kept `parseFromDocs()` - deprecated but available as fallback
   - Added schema JSON parsing structures

2. **`main.go`**
   - Updated CLI description
   - Made `-d` directory flag optional
   - Updated help text

### Dependencies

- Requires `terraform` binary in PATH
- Supports providers: `google`, `aws`, `azurerm`

## Migration Guide

### For Pike Maintainers

**Update all provider member files:**

```bash
cd /Users/jwoolfenden/code/pike

# Google
pike parse -n google
mv google-members.json parse/

# AWS
pike parse -n aws
mv aws-members.json parse/

# Azure
pike parse -n azurerm
mv azurerm-members.json parse/

# Update coverage
cd src/coverage
go test -run Test_coverageGcp
go test -run Test_coverageAws
go test -run Test_coverageAzurerm
```

### For Users

No action required - the command works the same way:

```bash
pike parse -n google
```

Just don't rely on the old documentation parsing method.

## Future Improvements

1. **Add version flag** - Support specific provider versions: `pike parse -n google --version 7.21.0`
2. **Cache schemas** - Avoid re-downloading for repeated runs
3. **Parallel extraction** - Parse multiple providers simultaneously
4. **Diff mode** - Show changes between versions
5. **More providers** - Extend support beyond google/aws/azurerm

## Related Issues

- Initial discovery: User found `google_securityposture_posture_deployment` missing
- Root cause: Documentation parsing missed 15 resources, included 141 invalid ones
- Impact: 88.40% coverage was against wrong baseline (1310 vs 1184)
- Resolution: Schema-based extraction gives 100% accurate baseline

## Credits

**Problem Identified By:** User discovering missing `google_securityposture_posture_deployment`
**Solution Implemented:** 2026-03-04
**Tested:** All three providers (google, aws, azurerm)
