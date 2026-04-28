resource "azurerm_network_manager_routing_rule" "pike_gen" {
  name               = "example-routing-rule"
  rule_collection_id = "pike"
  description        = "example routing rule"

  destination {
    type    = "AddressPrefix"
    address = "10.0.0.0/24"
  }

  next_hop {
    type = "VirtualNetworkGateway"
  }
}
