data "azurerm_mssql_elasticpool" "pike_gen" {
  name                = "mssqlelasticpoolname"
  resource_group_name = "example-resources"
  server_name         = "example-sql-server"
}
