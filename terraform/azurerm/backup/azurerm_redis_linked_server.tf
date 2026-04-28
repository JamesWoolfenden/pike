resource "azurerm_redis_linked_server" "pike_gen" {
  target_redis_cache_name     = "pike"
  resource_group_name         = "pike"
  linked_redis_cache_id       = "pike"
  linked_redis_cache_location = "pike"
  server_role                 = "Secondary"
}
