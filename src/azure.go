package pike

// GetAZUREPermissions for GCP resources.
func GetAZUREPermissions(result ResourceV2) ([]string, error) {
	var (
		err         error
		Permissions []string
	)

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

// GetAZUREResourcePermissions looks up permissions required for resources.
func GetAZUREResourcePermissions(result ResourceV2) ([]string, error) {
	temp := AzureLookup(result.Name)

	var (
		Permissions []string
		err         error
	)

	if temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes, result.Name)
	} else {
		//goland:noinspection GoLinter
		return nil, &notImplementedResourceError{result.Name}
	}

	return Permissions, err
}

func AzureLookup(name string) interface{} {
	TFLookup := map[string]interface{}{
		"azurerm_api_management":                       azurermAPIManagement,
		"azurerm_app_configuration":                    azurermAppConfiguration,
		"azurerm_app_service":                          azurermAppService,
		"azurerm_app_service_plan":                     azurermAppServicePlan,
		"azurerm_cognitive_account":                    azurermCognitiveAccount,
		"azurerm_container_registry":                   azurermContainerRegistry,
		"azurerm_cosmosdb_account":                     azurermCosmosdbAccount,
		"azurerm_cosmosdb_table":                       azurermCosmosdbTable,
		"azurerm_disk_encryption_set":                  azurermDiskEncryptionSet,
		"azurerm_dns_zone":                             azurermDNSZone,
		"azurerm_key_vault":                            azurermKeyVault,
		"azurerm_key_vault_access_policy":              azurermKeyVaultAccessPolicy,
		"azurerm_key_vault_key":                        azurermKeyVaultKey,
		"azurerm_linux_virtual_machine":                azurermVirtualMachine,
		"azurerm_linux_virtual_machine_scale_set":      azurermLinuxVirtualMachineScaleSet,
		"azurerm_log_analytics_solution":               azurermLogAnalyticsSolution,
		"azurerm_log_analytics_workspace":              azurermLogAnalyticsWorkspace,
		"azurerm_managed_disk":                         azurermManagedDisk,
		"azurerm_management_group":                     azurermManagementGroup,
		"azurerm_mariadb_configuration":                azurermMariadbConfiguration,
		"azurerm_mariadb_database":                     azurermMariadbDatabase,
		"azurerm_mariadb_firewall_rule":                azurermMariadbFirewallRule,
		"azurerm_mariadb_server":                       azurermMariadbServer,
		"azurerm_network_interface":                    azurermNetworkInterface,
		"azurerm_network_security_group":               azurermNetworkSecurityGroup,
		"azurerm_network_security_rule":                azurermNetworkSecurityRule,
		"azurerm_network_watcher":                      azurermNetworkWatcher,
		"azurerm_network_watcher_flow_log":             azurermNetworkWatcherFlowLog,
		"azurerm_private_dns_zone":                     azurermPrivateDNSZone,
		"azurerm_private_endpoint":                     azurermPrivateEndpoint,
		"azurerm_redis_cache":                          azurermRedisCache,
		"azurerm_resource_group":                       azurermResourceGroup,
		"azurerm_role_assignment":                      azurermRoleAssignment,
		"azurerm_search_service":                       azurermSearchService,
		"azurerm_security_center_contact":              azurermSecurityCenterContact,
		"azurerm_security_center_setting":              azurermSecurityCenterSetting,
		"azurerm_security_center_workspace":            azurermSecurityCenterWorkspace,
		"azurerm_service_plan":                         azurermServicePlan,
		"azurerm_storage_account":                      azurermStorageAccount,
		"azurerm_storage_account_customer_managed_key": azurermStorageAccountCustomerManagedKey,
		"azurerm_storage_account_network_rules":        azurermStorageAccountNetworkRules,
		"azurerm_storage_container":                    azurermStorageContainer,
		"azurerm_storage_sync":                         azurermStorageSync,
		"azurerm_storage_sync_group":                   azurermStorageSyncGroup,
		"azurerm_subnet":                               azurermSubnet,
		"azurerm_user_assigned_identity":               azurermUserAssignedIdentity,
		"azurerm_virtual_machine":                      azurermVirtualMachine,
		"azurerm_virtual_machine_scale_set":            azurermLinuxVirtualMachineScaleSet,
		"azurerm_virtual_network":                      azurermVirtualNetwork,
		"azurerm_virtual_network_peering":              azurermVirtualNetworkPeering,
		"azurerm_web_pubsub":                           azurermWebPubsub,
		"azurerm_windows_virtual_machine":              azurermVirtualMachine,
		"azurerm_windows_virtual_machine_scale_set":    azurermLinuxVirtualMachineScaleSet,
	}

	return TFLookup[name]
}
