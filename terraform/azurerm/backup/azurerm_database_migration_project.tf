resource "azurerm_database_migration_project" "pike_gen" {
  name                = "example-dbms-project"
  service_name        = "pike"
  resource_group_name = "pike"
  location            = "pike"
  source_platform     = "SQL"
  target_platform     = "SQLDB"
}
