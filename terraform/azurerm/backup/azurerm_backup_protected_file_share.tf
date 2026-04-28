resource "azurerm_backup_protected_file_share" "pike_gen" {
  resource_group_name       = "pike"
  recovery_vault_name       = "pike"
  source_storage_account_id = "pike"
  source_file_share_name    = "pike"
  backup_policy_id          = "pike"
}
