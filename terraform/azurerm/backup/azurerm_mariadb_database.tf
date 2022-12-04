resource "azurerm_mariadb_database" "pike" {
  name                = "mariadb_database"
  resource_group_name = "pike"
  server_name         = "mariadb-svr-jgw"
  charset             = "utf8mb4"
  collation           = "utf8mb4_unicode_520_ci"
}
