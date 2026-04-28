resource "azurerm_route_server_bgp_connection" "pike_gen" {
  name            = "example-rs-bgpconnection"
  route_server_id = "pike"
  peer_asn        = 65501
  peer_ip         = "169.254.21.5"
}
