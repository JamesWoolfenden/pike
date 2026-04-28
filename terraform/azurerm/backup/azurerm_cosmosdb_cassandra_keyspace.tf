resource "azurerm_cosmosdb_cassandra_keyspace" "pike_gen" {
  name                = "tfex-cosmos-cassandra-keyspace"
  resource_group_name = "pike"
  account_name        = "pike"
  throughput          = 400
}
