resource "azurerm_network_security_perimeter" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "West Europe"
}
