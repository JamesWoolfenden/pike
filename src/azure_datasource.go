package pike

import "github.com/rs/zerolog/log"

// GetAZUREDataPermissions gets permissions required for datasources
func GetAZUREDataPermissions(result ResourceV2) ([]string, error) {

	TFLookup := map[string]interface{}{
		"azurerm_resource_group": dataAzurermResourceGroup,
		"azurerm_client_config":  placeholder,
		"azurerm_key_vault":      dataAzurermKeyVault,
	}

	var Permissions []string
	var err error
	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}
