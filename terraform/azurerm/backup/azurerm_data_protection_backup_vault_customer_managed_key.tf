resource "azurerm_data_protection_backup_vault_customer_managed_key" "pike_gen" {
  data_protection_backup_vault_id = "pike"
  key_vault_key_id                = "pike"
}
