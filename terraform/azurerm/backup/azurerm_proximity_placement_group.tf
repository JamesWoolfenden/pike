resource "azurerm_proximity_placement_group" "pike" {
  name                = "pike-ppg"
  location            = "uksouth"
  resource_group_name = "pike"

  tags = {
    pike = "permissions"
  }
}
