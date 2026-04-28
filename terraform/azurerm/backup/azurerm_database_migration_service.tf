resource "azurerm_database_migration_service" "pike_gen" {
  name                = "example-dms"
  location            = "pike"
  resource_group_name = "pike"
  subnet_id           = "pike"
  sku_name            = "Standard_1vCores"
}
