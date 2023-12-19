data "azurerm_redis_enterprise_database" "pike" {
  cluster_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.Cache/redisEnterprise/pike"
  name       = "pike"
}
