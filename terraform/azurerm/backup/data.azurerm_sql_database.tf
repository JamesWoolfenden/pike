data "azurerm_sql_database" "pike" {
  server_name         = "pike"
  resource_group_name = "pike"
  name                = "pike"
}
