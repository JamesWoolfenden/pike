resource "azurerm_postgresql_active_directory_administrator" "pike_gen" {
  server_name         = "pike"
  resource_group_name = "pike"
  login               = "sqladmin"
  tenant_id           = "pike"
  object_id           = "pike"
}
