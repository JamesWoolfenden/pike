data "azurerm_mssql_server" "pike_gen" {
  name                = "existingMsSqlServer"
  resource_group_name = "existingResGroup"
}
