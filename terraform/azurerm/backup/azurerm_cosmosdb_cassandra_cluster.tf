resource "azurerm_cosmosdb_cassandra_cluster" "pike_gen" {
  name                           = "example-cluster"
  resource_group_name            = "pike"
  location                       = "pike"
  delegated_management_subnet_id = "pike"

  depends_on = ["pike"]
}
