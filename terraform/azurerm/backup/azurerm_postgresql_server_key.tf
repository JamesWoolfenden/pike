resource "azurerm_postgresql_server_key" "pike_gen" {
  server_id        = "pike"
  key_vault_key_id = "pike"
}
