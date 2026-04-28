resource "azurerm_postgresql_flexible_server_backup" "pike_gen" {
  name      = "example-pfsb"
  server_id = "pike"
}
