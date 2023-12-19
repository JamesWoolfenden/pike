data "azurerm_backup_policy_file_share" "pike" {
  resource_group_name = "pike"
  name                = "pike"
  recovery_vault_name = "pike"
}
