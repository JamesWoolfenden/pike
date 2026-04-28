resource "azurerm_mssql_firewall_rule" "pike_gen" {
  name             = "FirewallRule1"
  server_id        = "pike"
  start_ip_address = "10.0.17.62"
  end_ip_address   = "10.0.17.62"
}
