data "azurerm_database_migration_project" "pike_gen" {
  name                = "example-dbms-project"
  resource_group_name = "example-rg"
  service_name        = "example-dbms"
}
