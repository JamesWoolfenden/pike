resource "azurerm_postgresql_flexible_server_database" "pike_gen" {
  name      = "exampledb"
  server_id = "pike"
  collation = "en_US.utf8"
  charset   = "UTF8"

  # prevent the possibility of accidental data loss
  lifecycle {
    prevent_destroy = true
  }
}
