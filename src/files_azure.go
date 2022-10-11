package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/azure/resource/resourcegroups/azurerm_resource_group.json
var azurermResourceGroup []byte

//go:embed mapping/azure/resource/serverfarms/azurerm_service_plan.json
var azurermServicePlan []byte
