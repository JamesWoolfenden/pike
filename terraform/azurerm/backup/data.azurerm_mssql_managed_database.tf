data "azurerm_mssql_managed_database" "pike_gen" {
  name                  = "example"
  resource_group_name   = "pike"
  managed_instance_name = "pike"
}
