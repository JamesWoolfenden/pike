resource "azurerm_spring_cloud_app_redis_association" "pike_gen" {
  name                = "example-bind"
  spring_cloud_app_id = "pike"
  redis_cache_id      = "pike"
  ssl_enabled         = true
}
