data "azurerm_postgresql_server" "pike_gen" {
  name                = "postgresql-server-1"
  resource_group_name = "api-rg-pro"
}
