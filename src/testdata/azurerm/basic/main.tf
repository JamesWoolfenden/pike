resource "azurerm_resource_group" "pike" {
  name     = "pike"
  location = "uksouth"
  tags     = { pike = "permissions" }
}
