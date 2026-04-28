resource "azurerm_mysql_flexible_database" "pike_gen" {
  name                = "exampledb"
  resource_group_name = "pike"
  server_name         = "pike"
  charset             = "utf8"
  collation           = "utf8_unicode_ci"
}
