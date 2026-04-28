resource "azurerm_cosmosdb_mongo_collection" "pike_gen" {
  name                = "tfex-cosmos-mongo-db"
  resource_group_name = "pike"
  account_name        = "pike"
  database_name       = "pike"

  default_ttl_seconds = "777"
  shard_key           = "uniqueKey"
  throughput          = 400

  index {
    keys   = ["_id"]
    unique = true
  }
}
