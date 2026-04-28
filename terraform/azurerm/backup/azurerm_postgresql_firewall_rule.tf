resource "azurerm_postgresql_firewall_rule" "pike_gen" {
  name                = "office"
  resource_group_name = "pike"
  server_name         = "pike"
  start_ip_address    = "40.112.8.12"
  end_ip_address      = "40.112.8.12"
}
