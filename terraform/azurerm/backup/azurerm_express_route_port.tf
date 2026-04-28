resource "azurerm_express_route_port" "pike_gen" {
  name                = "port1"
  resource_group_name = "pike"
  location            = "pike"
  peering_location    = "Airtel-Chennai-CLS"
  bandwidth_in_gbps   = 10
  encapsulation       = "Dot1Q"
}
