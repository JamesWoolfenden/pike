resource "azurerm_mariadb_configuration" "pike" {
  name                = "interactive_timeout"
  resource_group_name = "pike"
  server_name         = "mariadb-svr-jgw"
  value               = "600"
}
