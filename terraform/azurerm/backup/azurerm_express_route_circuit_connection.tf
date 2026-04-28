resource "azurerm_express_route_circuit_connection" "pike_gen" {
  name                = "example-ercircuitconnection"
  peering_id          = "pike"
  peer_peering_id     = "pike"
  address_prefix_ipv4 = "192.169.9.0/29"
  authorization_key   = "846a1918-b7a2-4917-b43c-8c4cdaee006a"
}
