resource "azurerm_vpn_site" "pike_gen" {
  name                = "site1"
  resource_group_name = "pike"
  location            = "pike"
  virtual_wan_id      = "pike"
  address_cidrs       = ["10.0.0.0/24"]

  link {
    name       = "link1"
    ip_address = "10.0.0.1"
  }
}
