#!/usr/bin/env python3
"""
Validate existing Azure datasource mappings against best practices and Azure permissions data.

This script checks existing datasource JSON files for:
1. Structural correctness (empty apply/destroy/modify)
2. Permission validity (only read-like permissions in plan)
3. Comparison against azure_permissions.json (if available)
"""

import json
import sys
from pathlib import Path
from collections import defaultdict


class DatasourceValidator:
    def __init__(self, datasources_dir, permissions_file=None):
        self.datasources_dir = Path(datasources_dir)
        self.permissions_file = Path(permissions_file) if permissions_file else None
        self.azure_permissions = None
        self.issues = defaultdict(list)
        self.stats = {
            'total': 0,
            'valid': 0,
            'warnings': 0,
            'errors': 0
        }

        if self.permissions_file and self.permissions_file.exists():
            with open(self.permissions_file, 'r') as f:
                self.azure_permissions = json.load(f)
            print(f"✓ Loaded Azure permissions data from {self.permissions_file}\n")
        else:
            print(f"⚠ No Azure permissions data available - validation will be limited\n")

    def validate_structure(self, resource_name, data):
        """Validate basic datasource structure."""
        errors = []
        warnings = []

        if not isinstance(data, list) or len(data) != 1:
            errors.append("Must be an array with exactly one object")
            return errors, warnings

        obj = data[0]

        # Check required fields
        required_fields = ['apply', 'destroy', 'modify', 'plan', 'attributes']
        for field in required_fields:
            if field not in obj:
                errors.append(f"Missing required field: {field}")

        # Check that apply/destroy/modify are empty
        for field in ['apply', 'destroy', 'modify']:
            if field in obj and obj[field]:
                errors.append(f"Datasource must have empty '{field}' array, found {len(obj[field])} items")

        # Check attributes
        if 'attributes' in obj:
            if not isinstance(obj['attributes'], dict):
                errors.append("'attributes' must be an object")
            elif 'tags' not in obj['attributes']:
                warnings.append("'attributes' should have 'tags' field")

        return errors, warnings

    def validate_permissions(self, resource_name, data):
        """Validate that permissions are appropriate for a datasource."""
        warnings = []
        errors = []

        if not isinstance(data, list) or len(data) == 0:
            return errors, warnings

        obj = data[0]
        plan_perms = obj.get('plan', [])

        # Check for write/delete permissions (should never be in datasources)
        for perm in plan_perms:
            perm_lower = perm.lower()
            if '/write' in perm_lower:
                errors.append(f"Write permission in plan: {perm}")
            elif '/delete' in perm_lower:
                errors.append(f"Delete permission in plan: {perm}")

        # Warn about unusual permission counts
        if len(plan_perms) == 0:
            warnings.append("No permissions in plan - datasource may not work")
        elif len(plan_perms) > 10:
            warnings.append(f"Many permissions ({len(plan_perms)}) - verify this is correct")

        return errors, warnings

    def compare_with_azure_permissions(self, resource_name, data):
        """Compare datasource permissions against Azure permissions data."""
        if not self.azure_permissions:
            return [], []

        errors = []
        warnings = []

        obj = data[0]
        plan_perms = set(obj.get('plan', []))

        # Try to map resource name to Azure resource type
        # This is heuristic-based matching
        resource_part = resource_name.replace('azurerm_', '').replace('_', '')

        potential_matches = []

        for provider, resources in self.azure_permissions.items():
            for resource_type, perms in resources.items():
                resource_type_normalized = resource_type.lower().replace('/', '').replace('_', '')
                if resource_type_normalized in resource_part or resource_part in resource_type_normalized:
                    potential_matches.append({
                        'provider': provider,
                        'resource_type': resource_type,
                        'perms': perms
                    })

        if not potential_matches:
            warnings.append(f"Could not find matching Azure resource type for '{resource_name}'")
            return errors, warnings

        # For now, just use the first match as a reference
        # In the future, could be smarter about this
        match = potential_matches[0]
        expected_read = set(match['perms'].get('read', []))
        expected_action = set(match['perms'].get('action', []))
        expected_write = set(match['perms'].get('write', []))
        expected_delete = set(match['perms'].get('delete', []))

        # Check if any plan permissions are write/delete from Azure data
        for perm in plan_perms:
            if perm in expected_write:
                errors.append(f"Contains write permission: {perm}")
            if perm in expected_delete:
                errors.append(f"Contains delete permission: {perm}")

        # Informational: show what could be available
        all_read_like = expected_read | expected_action
        missing = all_read_like - plan_perms
        extra = plan_perms - all_read_like - expected_write - expected_delete

        if missing and len(plan_perms) < len(expected_read):
            warnings.append(f"May be missing {len(missing)} read permissions from {match['provider']}/{match['resource_type']}")

        if extra:
            warnings.append(f"Has {len(extra)} permissions not in Azure data for {match['provider']}/{match['resource_type']}")

        if len(potential_matches) > 1:
            warnings.append(f"Multiple possible Azure matches ({len(potential_matches)}) - may need manual review")

        return errors, warnings

    def validate_file(self, filepath):
        """Validate a single datasource JSON file."""
        resource_name = filepath.stem
        self.stats['total'] += 1

        try:
            with open(filepath, 'r') as f:
                data = json.load(f)
        except json.JSONDecodeError as e:
            self.issues[resource_name].append(('ERROR', f"Invalid JSON: {e}"))
            self.stats['errors'] += 1
            return

        # Run validations
        errors, warnings = self.validate_structure(resource_name, data)
        for error in errors:
            self.issues[resource_name].append(('ERROR', error))
            self.stats['errors'] += 1

        for warning in warnings:
            self.issues[resource_name].append(('WARNING', warning))
            self.stats['warnings'] += 1

        errors, warnings = self.validate_permissions(resource_name, data)
        for error in errors:
            self.issues[resource_name].append(('ERROR', error))
            self.stats['errors'] += 1

        for warning in warnings:
            self.issues[resource_name].append(('WARNING', warning))
            self.stats['warnings'] += 1

        if self.azure_permissions:
            errors, warnings = self.compare_with_azure_permissions(resource_name, data)
            for error in errors:
                self.issues[resource_name].append(('ERROR', error))
                self.stats['errors'] += 1

            for warning in warnings:
                self.issues[resource_name].append(('WARNING', warning))
                self.stats['warnings'] += 1

        if not self.issues[resource_name]:
            self.stats['valid'] += 1

    def validate_all(self):
        """Validate all datasource files."""
        # Find all JSON files in datasources directory
        json_files = list(self.datasources_dir.rglob('*.json'))

        if not json_files:
            print(f"No JSON files found in {self.datasources_dir}")
            return

        print(f"Validating {len(json_files)} datasource mappings...\n")

        for filepath in sorted(json_files):
            self.validate_file(filepath)

    def print_report(self):
        """Print validation report."""
        print("=" * 80)
        print("VALIDATION REPORT")
        print("=" * 80)

        # Print issues
        if self.issues:
            for resource_name in sorted(self.issues.keys()):
                issues_list = self.issues[resource_name]
                print(f"\n{resource_name}:")
                for severity, message in issues_list:
                    color = '\033[91m' if severity == 'ERROR' else '\033[93m'
                    reset = '\033[0m'
                    print(f"  {color}[{severity}]{reset} {message}")
        else:
            print("\n✓ No issues found!")

        # Print statistics
        print("\n" + "=" * 80)
        print("STATISTICS")
        print("=" * 80)
        print(f"Total datasources:  {self.stats['total']}")
        print(f"Valid:              {self.stats['valid']}")
        print(f"With warnings:      {self.stats['total'] - self.stats['valid'] - (self.stats['errors'] // max(1, self.stats['total']))}")
        print(f"Total warnings:     {self.stats['warnings']}")
        print(f"Total errors:       {self.stats['errors']}")

        if self.stats['errors'] > 0:
            print("\n⚠ Datasources with errors should be reviewed and fixed")
        elif self.stats['warnings'] > 0:
            print("\n✓ No critical errors, but some warnings to review")
        else:
            print("\n✓ All datasources validated successfully!")

        return self.stats['errors']


def main():
    import argparse

    parser = argparse.ArgumentParser(
        description='Validate Azure datasource mappings'
    )
    parser.add_argument(
        '--datasources-dir',
        '-d',
        default='src/mapping/azurerm/data',
        help='Directory containing datasource JSON files (default: src/mapping/azurerm/data)'
    )
    parser.add_argument(
        '--permissions-file',
        '-p',
        default='terraform/azurerm/azure_permissions.json',
        help='Azure permissions JSON file (default: terraform/azurerm/azure_permissions.json)'
    )

    args = parser.parse_args()

    validator = DatasourceValidator(args.datasources_dir, args.permissions_file)
    validator.validate_all()
    errors = validator.print_report()

    sys.exit(1 if errors > 0 else 0)


if __name__ == '__main__':
    main()
