resource "azurerm_lb" "pike_gen" {
  name                = "TestLoadBalancer"
  location            = "pike"
  resource_group_name = "pike"

  frontend_ip_configuration {
    name                 = "PublicIPAddress"
    public_ip_address_id = "pike"
  }
}
