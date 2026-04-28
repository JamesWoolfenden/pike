resource "azurerm_data_protection_backup_instance_postgresql" "pike_gen" {
  name                                    = "example"
  location                                = "pike"
  vault_id                                = "pike"
  database_id                             = "pike"
  backup_policy_id                        = "pike"
  database_credential_key_vault_secret_id = "pike"
}
