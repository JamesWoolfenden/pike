resource "azurerm_mysql_flexible_server_active_directory_administrator" "pike_gen" {
  server_id   = "pike"
  identity_id = "pike"
  login       = "sqladmin"
  object_id   = "pike"
  tenant_id   = "pike"
}
