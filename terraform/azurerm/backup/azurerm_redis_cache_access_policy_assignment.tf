resource "azurerm_redis_cache_access_policy_assignment" "pike_gen" {
  name               = "example"
  redis_cache_id     = "pike"
  access_policy_name = "Data Contributor"
  object_id          = "pike"
  object_id_alias    = "ServicePrincipal"
}
