resource "azurerm_network_interface_application_gateway_backend_address_pool_association" "pike_gen" {
  network_interface_id    = "pike"
  ip_configuration_name   = "testconfiguration1"
  backend_address_pool_id = tolist("pike").0.id
}
