resource "azurerm_backup_container_storage_account" "pike_gen" {
  resource_group_name = "pike"
  recovery_vault_name = "pike"
  storage_account_id  = "pike"
}
