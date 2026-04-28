resource "azurerm_postgresql_flexible_server_active_directory_administrator" "pike_gen" {
  server_name         = "pike"
  resource_group_name = "pike"
  tenant_id           = "pike"
  object_id           = "pike"
  principal_name      = "pike"
  principal_type      = "ServicePrincipal"
}
