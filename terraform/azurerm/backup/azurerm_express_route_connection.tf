resource "azurerm_express_route_connection" "pike_gen" {
  name                             = "example-expressrouteconn"
  express_route_gateway_id         = "pike"
  express_route_circuit_peering_id = "pike"
}
