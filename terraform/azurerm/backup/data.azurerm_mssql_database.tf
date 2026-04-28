data "azurerm_mssql_database" "pike_gen" {
  name      = "example-mssql-db"
  server_id = "pike"
}
