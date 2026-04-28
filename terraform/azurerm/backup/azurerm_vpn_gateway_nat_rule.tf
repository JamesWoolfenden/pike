resource "azurerm_vpn_gateway_nat_rule" "pike_gen" {
  name           = "example-vpngatewaynatrule"
  vpn_gateway_id = "pike"

  external_mapping {
    address_space = "192.168.21.0/26"
  }

  internal_mapping {
    address_space = "10.4.0.0/26"
  }
}
