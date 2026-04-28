resource "azurerm_cosmosdb_cassandra_datacenter" "pike_gen" {
  name                           = "example-datacenter"
  location                       = "pike"
  cassandra_cluster_id           = "pike"
  delegated_management_subnet_id = "pike"
  node_count                     = 3
  disk_count                     = 4
  sku_name                       = "Standard_DS14_v2"
  availability_zones_enabled     = false
}
