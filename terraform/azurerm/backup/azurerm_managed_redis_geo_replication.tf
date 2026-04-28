resource "azurerm_managed_redis_geo_replication" "pike_gen" {
  managed_redis_id = "pike"

  linked_managed_redis_ids = [
    azurerm_managed_redis.amr2.id,
  ]
}
