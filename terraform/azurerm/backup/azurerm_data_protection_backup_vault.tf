resource "azurerm_data_protection_backup_vault" "pike_gen" {
  name                = "example-backup-vault"
  resource_group_name = "pike"
  location            = "pike"
  datastore_type      = "VaultStore"
  redundancy          = "LocallyRedundant"
}
