resource "azurerm_express_route_circuit_peering" "pike_gen" {
  peering_type                  = "MicrosoftPeering"
  express_route_circuit_name    = "pike"
  resource_group_name           = "pike"
  peer_asn                      = 100
  primary_peer_address_prefix   = "123.0.0.0/30"
  secondary_peer_address_prefix = "123.0.0.4/30"
  ipv4_enabled                  = true
  vlan_id                       = 300

  microsoft_peering_config {
    advertised_public_prefixes = ["123.1.0.0/24"]
  }

  ipv6 {
    primary_peer_address_prefix   = "2002:db01::/126"
    secondary_peer_address_prefix = "2003:db01::/126"
    enabled                       = true

    microsoft_peering {
      advertised_public_prefixes = ["2002:db01::/126"]
    }
  }
}
