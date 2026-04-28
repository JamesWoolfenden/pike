data "azurerm_express_route_circuit_peering" "pike_gen" {
  peering_type               = "example-peering"
  express_route_circuit_name = "example-expressroute"
  resource_group_name        = "example-resources"
}
