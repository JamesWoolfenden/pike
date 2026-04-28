resource "azurerm_virtual_wan" "pike_gen" {
  name                = "example-vwan"
  resource_group_name = "pike"
  location            = "pike"
}
