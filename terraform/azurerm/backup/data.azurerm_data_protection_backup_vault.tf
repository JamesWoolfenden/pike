data "azurerm_data_protection_backup_vault" "pike_gen" {
  name                = "existing-backup-vault"
  resource_group_name = "existing-resource-group"
}
