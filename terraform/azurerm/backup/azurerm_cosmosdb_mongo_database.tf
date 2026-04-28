resource "azurerm_cosmosdb_mongo_database" "pike_gen" {
  name                = "tfex-cosmos-mongo-db"
  resource_group_name = "pike"
  account_name        = "pike"
  throughput          = 400
}
