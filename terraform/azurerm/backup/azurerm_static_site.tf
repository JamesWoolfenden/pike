resource "azurerm_static_site" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
}
