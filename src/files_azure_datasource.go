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
