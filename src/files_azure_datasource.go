package pike

import (
	_ "embed" // required for embed
)

//go:embed  mapping/azure/data/resourcegroups/azurerm_resource_group.json
var dataAzurermResourceGroup []byte
