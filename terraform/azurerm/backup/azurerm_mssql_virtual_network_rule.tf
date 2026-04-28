resource "azurerm_mssql_virtual_network_rule" "pike_gen" {
  name      = "sql-vnet-rule"
  server_id = "pike"
  subnet_id = "pike"
}
