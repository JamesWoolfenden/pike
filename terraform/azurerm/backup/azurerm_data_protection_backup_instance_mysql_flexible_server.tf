resource "azurerm_data_protection_backup_instance_mysql_flexible_server" "pike_gen" {
  name             = "example-dbi"
  location         = "pike"
  vault_id         = "pike"
  server_id        = "pike"
  backup_policy_id = "pike"
}
