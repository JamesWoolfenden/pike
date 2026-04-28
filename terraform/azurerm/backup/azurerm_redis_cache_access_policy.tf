resource "azurerm_redis_cache_access_policy" "pike_gen" {
  name           = "example"
  redis_cache_id = "pike"
  permissions    = "+@read +@connection +cluster|info"
}
