resource "azurerm_postgresql_flexible_server_configuration" "pike_gen" {
  name      = "backslash_quote"
  server_id = "pike"
  value     = "on"
}
