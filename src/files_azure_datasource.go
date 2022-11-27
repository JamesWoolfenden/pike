package pike

import (
	_ "embed" // required for embed
)

//go:embed  mapping/azurerm/data/resourcegroups/azurerm_resource_group.json
var dataAzurermResourceGroup []byte

//go:embed  mapping/azurerm/data/keyvault/azurerm_key_vault.json
var dataAzurermKeyVault []byte

//go:embed  mapping/azurerm/data/keyvault/azurerm_key_vault_key.json
var dataAzurermKeyVaultKey []byte

//go:embed mapping/azurerm/data/network/azurerm_subnet.json
var dataAzurermSubnet []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_network.json
var dataAzurermVirtualNetwork []byte

//go:embed mapping/azurerm/data/resources/azurerm_subscription.json
var dataAzurermSubscription []byte
