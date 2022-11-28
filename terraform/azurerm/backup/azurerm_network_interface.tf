resource "azurerm_network_interface" "pike" {
  name                = "example-nic"
  location            = data.azurerm_resource_group.pike.location
  resource_group_name = data.azurerm_resource_group.pike.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = data.azurerm_subnet.pike.id
    private_ip_address_allocation = "Dynamic"
  }
}
