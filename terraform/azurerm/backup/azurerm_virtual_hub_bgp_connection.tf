resource "azurerm_virtual_hub_bgp_connection" "pike_gen" {
  name           = "example-vhub-bgpconnection"
  virtual_hub_id = "pike"
  peer_asn       = 65514
  peer_ip        = "169.254.21.5"

  depends_on = ["pike"]
}
