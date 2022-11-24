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
