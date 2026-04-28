resource "azurerm_point_to_site_vpn_gateway" "pike_gen" {
  name                        = "example-vpn-gateway"
  location                    = "pike"
  resource_group_name         = "pike"
  virtual_hub_id              = "pike"
  vpn_server_configuration_id = "pike"
  scale_unit                  = 1
  connection_configuration {
    name = "example-gateway-config"

    vpn_client_address_pool {
      address_prefixes = [
        "10.0.2.0/24"
      ]
    }
  }
}
