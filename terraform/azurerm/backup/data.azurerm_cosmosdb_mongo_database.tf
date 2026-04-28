data "azurerm_cosmosdb_mongo_database" "pike_gen" {
  name                = "test-cosmosdb-mongo-db"
  resource_group_name = "test-cosmosdb-account-rg"
  account_name        = "test-cosmosdb-account"
}
