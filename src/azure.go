package pike

import (
	"errors"
	"fmt"
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
		"azurerm_resource_group":                    azurermResourceGroup,
		"azurerm_service_plan":                      azurermServicePlan,
		"azurerm_key_vault":                         azurermKeyVault,
		"azurerm_cosmosdb_account":                  azurermCosmosdbAccount,
		"azurerm_cosmosdb_table":                    azurermCosmosdbTable,
		"azurerm_container_registry":                azurermContainerRegistry,
		"azurerm_user_assigned_identity":            azurermUserAssignedIdentity,
		"azurerm_managed_disk":                      azurermManagedDisk,
		"azurerm_subnet":                            azurermSubnet,
		"azurerm_virtual_network":                   azurermVirtualNetwork,
		"azurerm_linux_virtual_machine_scale_set":   azurermLinuxVirtualMachineScaleSet,
		"azurerm_windows_virtual_machine_scale_set": azurermLinuxVirtualMachineScaleSet,
		"azurerm_virtual_machine_scale_set":         azurermLinuxVirtualMachineScaleSet,
		"azurerm_linux_virtual_machine":             azurermVirtualMachine,
		"azurerm_windows_virtual_machine":           azurermVirtualMachine,
		"azurerm_virtual_machine":                   azurermVirtualMachine,
		"azurerm_network_interface":                 azurermNetworkInterface,
		"azurerm_management_group":                  azurermManagementGroup,
		"azurerm_network_security_group":            azurermNetworkSecurityGroup,
		"azurerm_network_security_rule":             azurermNetworkSecurityRule,
		"azurerm_network_watcher":                   azurermNetworkWatcher,
		"azurerm_log_analytics_workspace":           azurermLogAnalyticsWorkspace,
		"azurerm_network_watcher_flow_log":          azurermNetworkWatcherFlowLog,
		"azurerm_storage_account":                   azurermStorageAccount,
	}

	var Permissions []string
	var err error

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		message := fmt.Sprintf("%s not implemented", result.Name)
		//log.Print(message)
		return nil, errors.New(message)
	}

	return Permissions, err
}
