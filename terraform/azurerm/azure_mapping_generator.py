#!/usr/bin/env python3
"""
Azure Mapping Generator for Pike.

This script generates Pike mapping JSON files for Azure resources based on
the parsed Azure permissions data.

Input: azure_permissions.json (from parse_azure_permissions.py)
Output: JSON mapping files in src/mapping/azurerm/resource/

Usage:
    python azure_mapping_generator.py
    python azure_mapping_generator.py --input azure_permissions.json
    python azure_mapping_generator.py --dry-run  # Preview without writing files
"""

import json
import sys
import argparse
import re
from pathlib import Path
from collections import defaultdict
from typing import Dict, List, Set


# Mapping of Azure provider/resource to Terraform resource names
# Format: "Microsoft.Provider/resourceType" -> "azurerm_resource_name"
RESOURCE_NAME_MAPPINGS = {
    # Compute
    "Microsoft.Compute/virtualMachines": "azurerm_virtual_machine",
    "Microsoft.Compute/disks": "azurerm_managed_disk",
    "Microsoft.Compute/snapshots": "azurerm_snapshot",
    "Microsoft.Compute/availabilitySets": "azurerm_availability_set",
    "Microsoft.Compute/virtualMachineScaleSets": "azurerm_virtual_machine_scale_set",

    # Storage
    "Microsoft.Storage/storageAccounts": "azurerm_storage_account",
    "Microsoft.Storage/storageAccounts/blobServices": "azurerm_storage_blob",
    "Microsoft.Storage/storageAccounts/queueServices": "azurerm_storage_queue",
    "Microsoft.Storage/storageAccounts/tableServices": "azurerm_storage_table",

    # Network
    "Microsoft.Network/virtualNetworks": "azurerm_virtual_network",
    "Microsoft.Network/networkInterfaces": "azurerm_network_interface",
    "Microsoft.Network/publicIPAddresses": "azurerm_public_ip",
    "Microsoft.Network/loadBalancers": "azurerm_lb",
    "Microsoft.Network/applicationGateways": "azurerm_application_gateway",
    "Microsoft.Network/networkSecurityGroups": "azurerm_network_security_group",
    "Microsoft.Network/routeTables": "azurerm_route_table",

    # Container
    "Microsoft.ContainerRegistry/registries": "azurerm_container_registry",
    "Microsoft.ContainerService/managedClusters": "azurerm_kubernetes_cluster",
    "Microsoft.ContainerInstance/containerGroups": "azurerm_container_group",

    # Database
    "Microsoft.Sql/servers": "azurerm_mssql_server",
    "Microsoft.Sql/servers/databases": "azurerm_mssql_database",
    "Microsoft.DBforPostgreSQL/servers": "azurerm_postgresql_server",
    "Microsoft.DBforMySQL/servers": "azurerm_mysql_server",
    "Microsoft.DocumentDB/databaseAccounts": "azurerm_cosmosdb_account",

    # KeyVault
    "Microsoft.KeyVault/vaults": "azurerm_key_vault",
    "Microsoft.KeyVault/vaults/secrets": "azurerm_key_vault_secret",
    "Microsoft.KeyVault/vaults/keys": "azurerm_key_vault_key",

    # Web
    "Microsoft.Web/sites": "azurerm_app_service",
    "Microsoft.Web/serverFarms": "azurerm_app_service_plan",
    "Microsoft.Web/staticSites": "azurerm_static_site",

    # Management
    "Microsoft.ManagedIdentity/userAssignedIdentities": "azurerm_user_assigned_identity",
    "Microsoft.Resources/resourceGroups": "azurerm_resource_group",

    # Monitoring
    "Microsoft.OperationalInsights/workspaces": "azurerm_log_analytics_workspace",
    "Microsoft.Insights/components": "azurerm_application_insights",

    # Cache
    "Microsoft.Cache/Redis": "azurerm_redis_cache",

    # Service Bus
    "Microsoft.ServiceBus/namespaces": "azurerm_servicebus_namespace",
    "Microsoft.ServiceBus/namespaces/queues": "azurerm_servicebus_queue",
    "Microsoft.ServiceBus/namespaces/topics": "azurerm_servicebus_topic",

    # Event Hub
    "Microsoft.EventHub/namespaces": "azurerm_eventhub_namespace",
    "Microsoft.EventHub/namespaces/eventhubs": "azurerm_eventhub",

    # Cognitive Services
    "Microsoft.CognitiveServices/accounts": "azurerm_cognitive_account",

    # API Management
    "Microsoft.ApiManagement/service": "azurerm_api_management",
}


def convert_azure_resource_to_terraform(provider: str, resource_path: str) -> str:
    """
    Convert Azure resource path to Terraform resource name.

    Args:
        provider: e.g., "Microsoft.Compute"
        resource_path: e.g., "virtualMachines"

    Returns:
        Terraform resource name, e.g., "azurerm_virtual_machine"
    """
    full_path = f"{provider}/{resource_path}"

    # Check if we have a specific mapping
    if full_path in RESOURCE_NAME_MAPPINGS:
        return RESOURCE_NAME_MAPPINGS[full_path]

    # Fallback: generate a name based on convention
    # Microsoft.Compute/virtualMachines -> azurerm_virtual_machine
    # Remove Microsoft. prefix
    provider_short = provider.replace("Microsoft.", "").lower()

    # Convert camelCase to snake_case
    resource_snake = re.sub(r'(?<!^)(?=[A-Z])', '_', resource_path).lower()

    return f"azurerm_{provider_short}_{resource_snake}"


def get_provider_directory(provider: str) -> str:
    """
    Get the directory name for a provider.

    Args:
        provider: e.g., "Microsoft.Compute"

    Returns:
        Directory name, e.g., "compute"
    """
    return provider.replace("Microsoft.", "").lower()


def generate_mapping(resource_name: str, permissions: Dict[str, List[str]]) -> List[Dict]:
    """
    Generate Pike mapping structure for a resource.

    Args:
        resource_name: Terraform resource name (e.g., "azurerm_virtual_machine")
        permissions: Dict with keys: read, write, delete, action

    Returns:
        List with single mapping dict
    """
    # Always include resource group read for Azure resources
    base_permissions = [
        "Microsoft.Resources/subscriptions/resourcegroups/read"
    ]

    # Build apply permissions (read + write + action)
    apply_perms = base_permissions.copy()
    apply_perms.extend(permissions.get('read', []))
    apply_perms.extend(permissions.get('write', []))
    apply_perms.extend(permissions.get('action', []))

    # Build destroy permissions
    destroy_perms = permissions.get('delete', [])

    # Build plan permissions (read only)
    plan_perms = permissions.get('read', [])

    # Remove duplicates and sort
    apply_perms = sorted(set(apply_perms))
    destroy_perms = sorted(set(destroy_perms))
    plan_perms = sorted(set(plan_perms))

    mapping = {
        "apply": apply_perms,
        "attributes": {
            "tags": []
        },
        "destroy": destroy_perms,
        "modify": [],
        "plan": plan_perms
    }

    return [mapping]


def main():
    parser = argparse.ArgumentParser(
        description='Generate Pike Azure resource mappings'
    )
    parser.add_argument(
        '--input',
        '-i',
        default='azure_permissions.json',
        help='Input JSON file from parse_azure_permissions.py (default: azure_permissions.json)'
    )
    parser.add_argument(
        '--output-dir',
        '-o',
        default='../../src/mapping/azurerm/resource',
        help='Output directory for mapping files (default: ../../src/mapping/azurerm/resource)'
    )
    parser.add_argument(
        '--dry-run',
        action='store_true',
        help='Preview what would be generated without writing files'
    )
    parser.add_argument(
        '--filter',
        help='Only process resources matching this provider (e.g., Microsoft.Compute)'
    )

    args = parser.parse_args()

    input_file = Path(args.input)
    output_base = Path(__file__).parent / args.output_dir

    if not input_file.exists():
        print(f"ERROR: {input_file} not found!")
        print("\nPlease run parse_azure_permissions.py first:")
        print("  python parse_azure_permissions.py")
        sys.exit(1)

    print(f"Reading {input_file}...")
    with open(input_file, 'r', encoding='utf-8') as f:
        permissions_data = json.load(f)

    print(f"Processing {len(permissions_data)} providers...")

    stats = {
        'providers_processed': 0,
        'resources_generated': 0,
        'files_written': 0,
        'files_skipped': 0
    }

    for provider, resources in sorted(permissions_data.items()):
        # Filter by provider if specified
        if args.filter and not provider.startswith(args.filter):
            continue

        provider_dir = get_provider_directory(provider)
        output_dir = output_base / provider_dir

        print(f"\n[{provider}]")
        print(f"  Output directory: {output_dir}")

        if not args.dry_run:
            output_dir.mkdir(parents=True, exist_ok=True)

        stats['providers_processed'] += 1

        for resource_path, perms in sorted(resources.items()):
            # Convert to Terraform resource name
            tf_resource = convert_azure_resource_to_terraform(provider, resource_path)

            # Generate mapping
            mapping = generate_mapping(tf_resource, perms)

            # Output file path - sanitize filename by replacing slashes with underscores
            safe_filename = tf_resource.replace('/', '_')
            output_file = output_dir / f"{safe_filename}.json"

            # Check if file already exists
            if output_file.exists():
                print(f"  ⊗ {tf_resource}.json (already exists, skipping)")
                stats['files_skipped'] += 1
                continue

            if args.dry_run:
                print(f"  → {tf_resource}.json (would create)")
                print(f"     - Apply permissions: {len(mapping[0]['apply'])}")
                print(f"     - Destroy permissions: {len(mapping[0]['destroy'])}")
                print(f"     - Plan permissions: {len(mapping[0]['plan'])}")
            else:
                # Write the file
                with open(output_file, 'w', encoding='utf-8') as f:
                    json.dump(mapping, f, indent=2, ensure_ascii=False)
                    f.write('\n')

                print(f"  ✓ {tf_resource}.json")
                stats['files_written'] += 1

            stats['resources_generated'] += 1

    # Print summary
    print(f"\n{'='*60}")
    print("Summary:")
    print(f"{'='*60}")
    print(f"Providers processed: {stats['providers_processed']}")
    print(f"Resources generated: {stats['resources_generated']}")

    if args.dry_run:
        print(f"\nDRY RUN - No files were written")
        print(f"Remove --dry-run flag to write {stats['resources_generated']} files")
    else:
        print(f"Files written: {stats['files_written']}")
        print(f"Files skipped (already exist): {stats['files_skipped']}")
        print(f"\nOutput directory: {output_base.resolve()}")

    print(f"\n{'='*60}")
    print("Next steps:")
    print(f"{'='*60}")
    print("1. Review generated mapping files")
    print("2. Build pike: go build -o pike")
    print("3. Test with Azure terraform: ./pike scan -d <azure-tf-dir> -p azurerm")
    print("\nNote: Generated mappings may need manual refinement for:")
    print("  - Attribute-specific permissions")
    print("  - Complex nested resources")
    print("  - Runtime permissions (similar to GCP)")


if __name__ == '__main__':
    main()
