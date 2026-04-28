resource "azurerm_network_manager_ipam_pool_static_cidr" "pike_gen" {
  name             = "example-ipsc"
  ipam_pool_id     = "pike"
  address_prefixes = ["10.0.0.0/26", "10.0.0.128/27"]
}
