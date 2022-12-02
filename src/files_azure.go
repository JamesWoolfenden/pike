package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/azurerm/resource/resourcegroups/azurerm_resource_group.json
var azurermResourceGroup []byte

//go:embed mapping/azurerm/resource/serverfarms/azurerm_service_plan.json
var azurermServicePlan []byte

//go:embed mapping/azurerm/resource/keyvault/azurerm_key_vault.json
var azurermKeyVault []byte

//go:embed mapping/azurerm/resource/documentdb/azurerm_cosmosdb_account.json
var azurermCosmosdbAccount []byte

//go:embed mapping/azurerm/resource/documentdb/azurerm_cosmosdb_table.json
var azurermCosmosdbTable []byte

//go:embed mapping/azurerm/resource/containerregistry/azurerm_container_registry.json
var azurermContainerRegistry []byte

//go:embed mapping/azurerm/resource/managedidentity/azurerm_user_assigned_identity.json
var azurermUserAssignedIdentity []byte

//go:embed mapping/azurerm/resource/compute/azurerm_managed_disk.json
var azurermManagedDisk []byte

//go:embed mapping/azurerm/resource/network/azurerm_subnet.json
var azurermSubnet []byte

//go:embed mapping/azurerm/resource/network/azurerm_virtual_network.json
var azurermVirtualNetwork []byte

//go:embed mapping/azurerm/resource/compute/azurerm_linux_virtual_machine_scale_set.json
var azurermLinuxVirtualMachineScaleSet []byte

//go:embed mapping/azurerm/resource/compute/azurerm_virtual_machine.json
var azurermVirtualMachine []byte

//go:embed mapping/azurerm/resource/network/azurerm_network_interface.json
var azurermNetworkInterface []byte

//go:embed mapping/azurerm/resource/management/azurerm_management_group.json
var azurermManagementGroup []byte

//go:embed mapping/azurerm/resource/network/azurerm_network_security_group.json
var azurermNetworkSecurityGroup []byte

//go:embed mapping/azurerm/resource/network/azurerm_network_security_rule.json
var azurermNetworkSecurityRule []byte

//go:embed mapping/azurerm/resource/network/azurerm_network_watcher.json
var azurermNetworkWatcher []byte

//go:embed mapping/azurerm/resource/operationalinsights/azurerm_log_analytics_workspace.json
var azurermLogAnalyticsWorkspace []byte

//go:embed mapping/azurerm/resource/network/azurerm_network_watcher_flow_log.json
var azurermNetworkWatcherFlowLog []byte

//go:embed mapping/azurerm/resource/storage/azurerm_storage_account.json
var azurermStorageAccount []byte

//go:embed mapping/azurerm/resource/network/azurerm_virtual_network_peering.json
var azurermVirtualNetworkPeering []byte

//go:embed mapping/azurerm/resource/network/azurerm_private_dns_zone.json
var azurermPrivateDNSZone []byte

//go:embed mapping/azurerm/resource/network/azurerm_dns_zone.json
var azurermDNSZone []byte

//go:embed mapping/azurerm/resource/storage/azurerm_storage_account_customer_managed_key.json
var azurermStorageAccountCustomerManagedKey []byte

//go:embed mapping/azurerm/resource/storage/azurerm_storage_account_network_rules.json
var azurermStorageAccountNetworkRules []byte

//go:embed mapping/azurerm/resource/storage/azurerm_storage_container.json
var azurermStorageContainer []byte

//go:embed mapping/azurerm/resource/network/azurerm_private_endpoint.json
var azurermPrivateEndpoint []byte

//go:embed mapping/azurerm/resource/cache/azurerm_redis_cache.json
var azurermRedisCache []byte

//go:embed mapping/azurerm/resource/signalrservice/azurerm_web_pubsub.json
var azurermWebPubsub []byte

//go:embed mapping/azurerm/resource/apimanagement/azurerm_api_management.json
var azurermAPIManagement []byte

//go:embed mapping/azurerm/resource/keyvault/azurerm_key_vault_access_policy.json
var azurermKeyVaultAccessPolicy []byte

//go:embed mapping/azurerm/resource/keyvault/azurerm_key_vault_key.json
var azurermKeyVaultKey []byte
