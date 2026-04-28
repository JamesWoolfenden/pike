resource "azurerm_lb_backend_address_pool_address" "pike_gen" {
  name                    = "example"
  backend_address_pool_id = "pike"
  virtual_network_id      = "pike"
  ip_address              = "10.0.0.1"
}
