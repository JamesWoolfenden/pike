package pike

import (
	_ "embed" // required for embed
)

//go:embed mapping/azurerm/data/resourcegroups/azurerm_resource_group.json
var dataAzurermResourceGroup []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault.json
var dataAzurermKeyVault []byte

//go:embed mapping/azurerm/data/keyvault/azurerm_key_vault_key.json
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

//go:embed mapping/azurerm/data/web/azurerm_app_service.json
var dataAzurermAppService []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_certificate.json
var dataAzurermAppServiceCertificate []byte

//go:embed mapping/azurerm/data/resources/azurerm_app_service_certificate_order.json
var dataAzurermAppServiceCertificateOrder []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_environment.json
var dataAzurermAppServiceEnvironment []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_environment_v3.json
var dataAzurermAppServiceEnvironmentV3 []byte

//go:embed mapping/azurerm/data/web/azurerm_app_service_plan.json
var dataAzurermAppServicePlan []byte

//go:embed mapping/azurerm/data/network/azurerm_public_ip.json
var dataAzurermPublicIP []byte

//go:embed mapping/azurerm/data/network/azurerm_public_ip_prefix.json
var dataAzurermPublicIPPrefix []byte

//go:embed mapping/azurerm/data/network/azurerm_public_ips.json
var dataAzurermPublicIps []byte

//go:embed mapping/azurerm/data/web/azurerm_windows_function_app.json
var dataAzurermWindowsFunctionApp []byte

//go:embed mapping/azurerm/data/web/azurerm_windows_web_app.json
var dataAzurermWindowsWebApp []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_blob.json
var dataAzurermStorageBlob []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_container.json
var dataAzurermStorageContainer []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_encryption_scope.json
var dataAzurermStorageEncryptionScope []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_management_policy.json
var dataAzurermStorageManagementPolicy []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_share.json
var dataAzurermStorageShare []byte

//go:embed mapping/azurerm/data/storagesync/azurerm_storage_sync.json
var dataAzurermStorageSync []byte

//go:embed mapping/azurerm/data/storagesync/azurerm_storage_sync_group.json
var dataAzurermStorageSyncGroup []byte

//go:embed mapping/azurerm/data/storage/azurerm_storage_table_entity.json
var dataAzurermStorageTableEntity []byte
