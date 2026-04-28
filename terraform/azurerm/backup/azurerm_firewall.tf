resource "azurerm_firewall" "pike_gen" {
  name                = "testfirewall"
  location            = "pike"
  resource_group_name = "pike"
  sku_name            = "AZFW_VNet"
  sku_tier            = "Standard"

  ip_configuration {
    name                 = "configuration"
    subnet_id            = "pike"
    public_ip_address_id = "pike"
  }
}
