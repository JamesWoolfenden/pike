resource "azurerm_network_manager_routing_configuration" "pike_gen" {
  name               = "example-routing-configuration"
  network_manager_id = "pike"
  description        = "example routing configuration"
}
