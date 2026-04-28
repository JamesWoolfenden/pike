resource "azurerm_postgresql_flexible_server_firewall_rule" "pike_gen" {
  name             = "example-fw"
  server_id        = "pike"
  start_ip_address = "122.122.0.0"
  end_ip_address   = "122.122.0.0"
}
