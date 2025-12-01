#!/usr/bin/env python3
"""
Parse Azure permissions from PowerShell export and convert to Pike format.

This script reads the JSON output from Get-AzProviderOperation and converts it
into a structured format that can be used by the Pike Azure mapping generator.

Input: azure_permissions_full.json (from export_azure_permissions.ps1)
Output: azure_permissions.json (structured by provider and resource type)
"""

import json
import sys
import argparse
from collections import defaultdict
from pathlib import Path


def parse_operation(operation_string):
    """
    Parse an Azure operation string into components.

    Example: "Microsoft.Storage/storageAccounts/read"
    Returns: {
        "provider": "Microsoft.Storage",
        "resource_path": "storageAccounts",
        "operation": "read"
    }
    """
    parts = operation_string.split('/')
    if len(parts) < 2:
        return None

    provider = parts[0]
    operation = parts[-1]
    resource_path = '/'.join(parts[1:-1]) if len(parts) > 2 else parts[1]

    return {
        "provider": provider,
        "resource_path": resource_path,
        "operation": operation,
        "full_permission": operation_string
    }


def categorize_permission(operation):
    """
    Categorize a permission as read, write, delete, or action.
    """
    op_lower = operation.lower()

    if op_lower == 'read':
        return 'read'
    elif op_lower == 'write':
        return 'write'
    elif op_lower == 'delete':
        return 'delete'
    elif op_lower.startswith('list') or op_lower.startswith('get'):
        return 'read'
    else:
        return 'action'


def main():
    parser = argparse.ArgumentParser(
        description='Parse Azure permissions from PowerShell export'
    )
    parser.add_argument(
        '--input',
        '-i',
        default='azure_permissions_full.json',
        help='Input JSON file (default: azure_permissions_full.json)'
    )
    parser.add_argument(
        '--output',
        '-o',
        default='azure_permissions.json',
        help='Output JSON file (default: azure_permissions.json)'
    )

    args = parser.parse_args()

    input_file = Path(args.input)
    output_file = Path(args.output)

    if not input_file.exists():
        print(f"ERROR: {input_file} not found!")
        print("Please run export_azure_permissions.ps1 first")
        sys.exit(1)

    print(f"Reading {input_file}...")

    try:
        with open(input_file, 'r', encoding='utf-8') as f:
            content = f.read().strip()
            if not content:
                print(f"ERROR: {input_file} is empty!")
                print("Please run export_azure_permissions.ps1 first to generate the data")
                print("\nSteps:")
                print("  1. Install Az.Resources: Install-Module -Name Az.Resources -Scope CurrentUser")
                print("  2. Login to Azure: Connect-AzAccount")
                print("  3. Run: pwsh export_azure_permissions.ps1")
                sys.exit(1)
            raw_data = json.loads(content)
    except json.JSONDecodeError as e:
        print(f"ERROR: {input_file} contains invalid JSON!")
        print(f"JSON Error: {e}")
        print("\nThe file may be corrupted. Please re-run export_azure_permissions.ps1")
        sys.exit(1)

    # Handle both single object and array formats
    if isinstance(raw_data, dict):
        operations_list = [raw_data]
    else:
        operations_list = raw_data

    print(f"Processing {len(operations_list)} operations...")

    # Structure: provider -> resource_type -> [permissions]
    permissions_by_provider = defaultdict(lambda: defaultdict(lambda: {
        'read': [],
        'write': [],
        'delete': [],
        'action': []
    }))

    # Track statistics
    stats = {
        'total_operations': 0,
        'providers': set(),
        'skipped': 0
    }

    for op_data in operations_list:
        # Handle PowerShell JSON format
        if isinstance(op_data, dict):
            operation_string = op_data.get('Operation') or op_data.get('operation')
        else:
            continue

        if not operation_string:
            stats['skipped'] += 1
            continue

        parsed = parse_operation(operation_string)
        if not parsed:
            stats['skipped'] += 1
            continue

        provider = parsed['provider']
        resource_path = parsed['resource_path']
        operation = parsed['operation']
        full_permission = parsed['full_permission']

        # Categorize the permission
        category = categorize_permission(operation)

        # Add to the appropriate category
        permissions_by_provider[provider][resource_path][category].append(full_permission)

        stats['total_operations'] += 1
        stats['providers'].add(provider)

    # Convert defaultdict to regular dict for JSON serialization
    output_data = {}
    for provider, resources in sorted(permissions_by_provider.items()):
        output_data[provider] = {}
        for resource_type, perms in sorted(resources.items()):
            # Sort all permission arrays
            output_data[provider][resource_type] = {
                'read': sorted(set(perms['read'])),
                'write': sorted(set(perms['write'])),
                'delete': sorted(set(perms['delete'])),
                'action': sorted(set(perms['action']))
            }

    # Write output
    print(f"\nWriting {output_file}...")
    with open(output_file, 'w', encoding='utf-8') as f:
        json.dump(output_data, f, indent=2, ensure_ascii=False)

    # Print statistics
    print(f"\n{'='*60}")
    print("Statistics:")
    print(f"{'='*60}")
    print(f"Total operations processed: {stats['total_operations']}")
    print(f"Unique providers found: {len(stats['providers'])}")
    print(f"Skipped operations: {stats['skipped']}")
    print(f"\nOutput written to: {output_file}")
    print(f"\nProviders found:")
    for provider in sorted(stats['providers']):
        resource_count = len(output_data[provider])
        print(f"  - {provider}: {resource_count} resource types")

    print(f"\n{'='*60}")
    print("Next steps:")
    print(f"{'='*60}")
    print("1. Review azure_permissions.json to verify the data looks correct")
    print("2. Update azure_mapping_generator.py to use this file")
    print("3. Test generating datasource mappings")


if __name__ == '__main__':
    main()
