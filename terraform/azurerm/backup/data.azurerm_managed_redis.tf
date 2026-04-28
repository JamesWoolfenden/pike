data "azurerm_managed_redis" "pike_gen" {
  name                = "example-managed-redis"
  resource_group_name = "example-resources"
}
