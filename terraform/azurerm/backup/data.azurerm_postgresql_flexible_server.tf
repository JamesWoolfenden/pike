data "azurerm_postgresql_flexible_server" "pike_gen" {
  name                = "existing-postgresql-fs"
  resource_group_name = "existing-postgresql-resgroup"
}
