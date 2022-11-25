resource "azurerm_subnet" "pike" {
  name                 = "internal"
  resource_group_name  = "pike"
  virtual_network_name = "example-network"
  address_prefixes     = ["10.0.2.0/24"]
}
