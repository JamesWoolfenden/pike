data "azurerm_managed_redis_access_policy_assignment" "pike_gen" {
  object_id           = "00000000-0000-0000-0000-000000000000"
  managed_redis_name  = "example-managedredis"
  resource_group_name = "example-resources"
}
