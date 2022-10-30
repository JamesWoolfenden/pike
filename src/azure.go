package pike

import (
	"github.com/rs/zerolog/log"
)

// GetAZUREPermissions for GCP resources
func GetAZUREPermissions(result ResourceV2) ([]string, error) {
	var err error
	var Permissions []string
	if result.TypeName == "resource" {
		Permissions, err = GetAZUREResourcePermissions(result)
		if err != nil {
			return Permissions, err
		}
	} else {
		Permissions, err = GetAZUREDataPermissions(result)
		if err != nil {
			return Permissions, err
		}
	}

	return Permissions, err
}

// GetAZUREResourcePermissions looks up permissions required for resources
func GetAZUREResourcePermissions(result ResourceV2) ([]string, error) {
	TFLookup := map[string]interface{}{
		"azurerm_resource_group":   azurermResourceGroup,
		"azurerm_service_plan":     azurermServicePlan,
		"azurerm_key_vault":        azurermKeyVault,
		"azurerm_cosmosdb_account": azureCosmosdbAccount,
		"azurerm_cosmosdb_table":   azureCosmosdbTable,
	}

	var Permissions []string
	var err error

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions, err
}
