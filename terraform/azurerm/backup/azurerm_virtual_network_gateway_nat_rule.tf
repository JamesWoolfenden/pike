resource "azurerm_virtual_network_gateway_nat_rule" "pike_gen" {
  name                       = "example-vnetgwnatrule"
  resource_group_name        = "pike"
  virtual_network_gateway_id = "pike"
  mode                       = "EgressSnat"
  type                       = "Dynamic"
  ip_configuration_id        = "pike"

  external_mapping {
    address_space = "10.2.0.0/26"
    port_range    = "200"
  }

  internal_mapping {
    address_space = "10.4.0.0/26"
    port_range    = "400"
  }
}
