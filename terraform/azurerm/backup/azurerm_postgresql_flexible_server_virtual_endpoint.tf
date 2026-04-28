resource "azurerm_postgresql_flexible_server_virtual_endpoint" "pike_gen" {
  name              = "example-endpoint-1"
  source_server_id  = "pike"
  replica_server_id = "pike"
  type              = "ReadWrite"
}
