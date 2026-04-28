data "azurerm_mssql_failover_group" "pike_gen" {
  name      = "example"
  server_id = "example-sql-server"
}
