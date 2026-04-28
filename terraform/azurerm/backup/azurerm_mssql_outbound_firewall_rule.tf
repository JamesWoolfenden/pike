resource "azurerm_mssql_outbound_firewall_rule" "pike_gen" {
  name      = "sqlexamplefdqn.database.windows.net"
  server_id = "pike"
}
