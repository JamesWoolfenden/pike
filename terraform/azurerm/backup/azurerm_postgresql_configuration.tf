resource "azurerm_postgresql_configuration" "pike_gen" {
  name                = "backslash_quote"
  resource_group_name = "pike"
  server_name         = "pike"
  value               = "on"
}
