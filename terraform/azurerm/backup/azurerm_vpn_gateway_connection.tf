resource "azurerm_vpn_gateway_connection" "pike_gen" {
  name               = "example"
  vpn_gateway_id     = "pike"
  remote_vpn_site_id = "pike"

  vpn_link {
    name             = "link1"
    vpn_site_link_id = "pike"
  }

  vpn_link {
    name             = "link2"
    vpn_site_link_id = "pike"
  }
}
