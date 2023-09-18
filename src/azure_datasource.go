package pike

import (
	"fmt"
)

// GetAZUREDataPermissions gets permissions required for datasources
func GetAZUREDataPermissions(result ResourceV2) ([]string, error) {
	temp := AzureDataLookup(result.Name)

	var (
		Permissions []string
		err         error
	)

	if temp != nil {
		Permissions, err = GetPermissionMap(temp.([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}

func AzureDataLookup(name string) interface{} {
	TFLookup := map[string]interface{}{
		"azurerm_resource_group":                             dataAzurermResourceGroup,
		"azurerm_client_config":                              placeholder,
		"azurerm_key_vault":                                  dataAzurermKeyVault,
		"azurerm_key_vault_key":                              dataAzurermKeyVaultKey,
		"azurerm_subnet":                                     dataAzurermSubnet,
		"azurerm_virtual_network":                            dataAzurermVirtualNetwork,
		"azurerm_subscription":                               dataAzurermSubscription,
		"azurerm_network_watcher":                            dataAzurermNetworkWatcher,
		"azurerm_storage_account":                            dataAzurermStorageAccount,
		"azurerm_key_vault_access_policy":                    placeholder,
		"azurerm_key_vault_certificate":                      dataAzurermKeyVaultCertificate,
		"azurerm_key_vault_certificate_data":                 dataAzurermKeyVaultCertificateData,
		"azurerm_key_vault_certificate_issuer":               dataAzurermKeyVaultCertificateIssuer,
		"azurerm_key_vault_certificates":                     dataAzurermKeyVaultCertificates,
		"azurerm_key_vault_encrypted_value":                  placeholder,
		"azurerm_key_vault_managed_hardware_security_module": dataAzurermKeyVaultManagedHardwareSecurityModule,
		"azurerm_key_vault_secret":                           dataAzurermKeyVaultSecret,
		"azurerm_key_vault_secrets":                          dataAzurermKeyVaultSecrets,
		"azurerm_app_service":                                dataAzurermAppService,
		"azurerm_app_service_certificate":                    dataAzurermAppServiceCertificate,
		"azurerm_app_service_certificate_order":              dataAzurermAppServiceCertificateOrder,
		"azurerm_app_service_environment":                    dataAzurermAppServiceEnvironment,
		"azurerm_app_service_environment_v3":                 dataAzurermAppServiceEnvironmentV3,
		"azurerm_app_service_plan":                           dataAzurermAppServicePlan,
		"azurerm_public_ip":                                  dataAzurermPublicIp,
		"azurerm_public_ip_prefix":                           dataAzurermPublicIpPrefix,
		"azurerm_public_ips":                                 dataAzurermPublicIps,
		"azurerm_windows_function_app":                       dataAzurermWindowsFunctionApp,
		"azurerm_windows_web_app":                            dataAzurermWindowsWebApp,
		"azurerm_storage_account_blob_container_sas":         placeholder,
		"azurerm_storage_account_sas":                        placeholder,
		"azurerm_storage_blob":                               dataAzurermStorageBlob,
		"azurerm_storage_container":                          dataAzurermStorageContainer,
		"azurerm_storage_encryption_scope":                   dataAzurermStorageEncryptionScope,
		"azurerm_storage_management_policy":                  dataAzurermStorageManagementPolicy,
		"azurerm_storage_share":                              dataAzurermStorageShare,
		"azurerm_storage_sync":                               dataAzurermStorageSync,
		"azurerm_storage_sync_group":                         dataAzurermStorageSyncGroup,
		"azurerm_storage_table_entity":                       dataAzurermStorageTableEntity,
	}

	return TFLookup[name]
}
