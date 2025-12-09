resource "azurerm_availability_set" "pike" {
  name                = "pike-avset"
  location            = "uksouth"
  resource_group_name = "pike"

  tags = {
    pike = "permissions"
  }
}
