resource "azurerm_postgresql_database" "pike_gen" {
  name                = "exampledb"
  resource_group_name = "pike"
  server_name         = "pike"
  charset             = "UTF8"
  collation           = "English_United States.1252"

  # prevent the possibility of accidental data loss
  lifecycle {
    prevent_destroy = true
  }
}
