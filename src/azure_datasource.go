package pike

import (
	"fmt"
)

// GetAZUREDataPermissions gets permissions required for datasources
func GetAZUREDataPermissions(result ResourceV2) ([]string, error) {
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
	}

	var (
		Permissions []string
		err         error
	)

	temp := TFLookup[result.Name]

	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}
