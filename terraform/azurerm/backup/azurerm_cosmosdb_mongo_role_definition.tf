resource "azurerm_cosmosdb_mongo_role_definition" "pike_gen" {
  cosmos_mongo_database_id = "pike"
  role_name                = "example-roledefinition"
}
