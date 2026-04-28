resource "azurerm_mssql_managed_instance_active_directory_administrator" "pike_gen" {
  managed_instance_id = "pike"
  login_username      = "msadmin"
  object_id           = "pike"
  tenant_id           = "pike"
}
