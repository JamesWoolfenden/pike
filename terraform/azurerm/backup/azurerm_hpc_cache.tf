resource "azurerm_hpc_cache" "pike_gen" {
  name                = "examplehpccache"
  resource_group_name = "pike"
  location            = "pike"
  cache_size_in_gb    = 3072
  subnet_id           = "pike"
  sku_name            = "Standard_2G"
}
