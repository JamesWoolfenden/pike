resource "azurerm_network_manager_ipam_pool" "pike_gen" {
  name               = "example-ipam-pool"
  location           = "West Europe"
  network_manager_id = "pike"
  display_name       = "example-pool"
  address_prefixes   = ["10.0.0.0/24"]
}
