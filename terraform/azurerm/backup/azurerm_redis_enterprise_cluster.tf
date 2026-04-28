resource "azurerm_redis_enterprise_cluster" "pike_gen" {
  name                = "example-redisenterprise"
  resource_group_name = "pike"
  location            = "pike"

  sku_name = "EnterpriseFlash_F300-3"
}
