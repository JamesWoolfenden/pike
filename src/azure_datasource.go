package pike

import (
	"fmt"
)

// GetAZUREDataPermissions gets permissions required for datasources
func GetAZUREDataPermissions(result ResourceV2) ([]string, error) {

	TFLookup := map[string]interface{}{
		"azurerm_resource_group":  dataAzurermResourceGroup,
		"azurerm_client_config":   placeholder,
		"azurerm_key_vault":       dataAzurermKeyVault,
		"azurerm_key_vault_key":   dataAzurermKeyVaultKey,
		"azurerm_subnet":          dataAzurermSubnet,
		"azurerm_virtual_network": dataAzurermVirtualNetwork,
		"azurerm_subscription":    dataAzurermSubscription,
		"azurerm_network_watcher": dataAzurermNetworkWatcher,
	}

	var Permissions []string
	var err error
	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}
