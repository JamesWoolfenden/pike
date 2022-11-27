resource "azurerm_network_security_group" "pike" {
  name                = "acceptanceTestSecurityGroup1"
  location            = "uksouth"
  resource_group_name = "pike"

  security_rule {
    name                       = "test1234"
    priority                   = 100
    direction                  = "Inbound"
    access                     = "Allow"
    protocol                   = "Tcp"
    source_port_range          = "*"
    destination_port_range     = "*"
    source_address_prefix      = "*"
    destination_address_prefix = "*"
  }

  tags = {
    pike = "permissions"
  }
}
