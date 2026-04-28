resource "azurerm_mysql_flexible_server_configuration" "pike_gen" {
  name                = "interactive_timeout"
  resource_group_name = "pike"
  server_name         = "pike"
  value               = "600"
}
