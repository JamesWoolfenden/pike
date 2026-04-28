resource "azurerm_api_management_redis_cache" "pike_gen" {
  name              = "example-Redis-Cache"
  api_management_id = "pike"
  description       = "Redis cache instances"
  redis_cache_id    = "pike"
  cache_location    = "East Us"
}
