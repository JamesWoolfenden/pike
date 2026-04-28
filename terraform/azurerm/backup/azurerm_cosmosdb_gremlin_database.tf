resource "azurerm_cosmosdb_gremlin_database" "pike_gen" {
  name                = "tfex-cosmos-gremlin-db"
  resource_group_name = "pike"
  account_name        = "pike"
  throughput          = 400
}
