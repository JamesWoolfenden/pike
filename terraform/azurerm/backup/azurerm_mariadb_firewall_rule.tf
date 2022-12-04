resource "azurerm_mariadb_firewall_rule" "pike" {
  name                = "test-rule"
  resource_group_name = "pike"
  server_name         = "mariadb-svr-jgw"
  start_ip_address    = "40.112.8.12"
  end_ip_address      = "40.112.8.12"
}
