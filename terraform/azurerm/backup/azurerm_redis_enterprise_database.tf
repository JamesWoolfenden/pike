resource "azurerm_redis_enterprise_database" "pike_gen" {
  name = "default"

  cluster_id        = "pike"
  client_protocol   = "Encrypted"
  clustering_policy = "EnterpriseCluster"
  eviction_policy   = "NoEviction"
  port              = 10000

  linked_database_id = [
    "${azurerm_redis_enterprise_cluster.example.id}/databases/default",
    "${azurerm_redis_enterprise_cluster.example1.id}/databases/default"
  ]

  linked_database_group_nickname = "tftestGeoGroup"
}
