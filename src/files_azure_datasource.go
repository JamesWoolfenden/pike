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

//go:embed mapping/azurerm/data/network/azurerm_subnet.json
var dataAzurermSubnet []byte

//go:embed mapping/azurerm/data/network/azurerm_virtual_network.json
var dataAzurermVirtualNetwork []byte

//go:embed mapping/azurerm/data/resources/azurerm_subscription.json
var dataAzurermSubscription []byte

//go:embed mapping/azurerm/data/network/azurerm_network_watcher.json
var dataAzurermNetworkWatcher []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_account.json
var dataAzurermStorageAccount []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificate.json
var dataAzurermKeyVaultCertificate []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificate_data.json
var dataAzurermKeyVaultCertificateData []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificate_issuer.json
var dataAzurermKeyVaultCertificateIssuer []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_certificates.json
var dataAzurermKeyVaultCertificates []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_managed_hardware_security_module.json
var dataAzurermKeyVaultManagedHardwareSecurityModule []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_secret.json
var dataAzurermKeyVaultSecret []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_secrets.json
var dataAzurermKeyVaultSecrets []byte
