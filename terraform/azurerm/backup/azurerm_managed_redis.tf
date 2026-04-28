resource "azurerm_managed_redis" "pike_gen" {
  name                = "example-managed-redis"
  resource_group_name = "pike"
  location            = "pike"
  sku_name            = "Balanced_B3"

  default_database {
    geo_replication_group_name = "myGeoGroup"
  }
}
