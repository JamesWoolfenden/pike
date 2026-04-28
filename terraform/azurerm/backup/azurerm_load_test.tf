resource "azurerm_load_test" "pike_gen" {
  location            = "pike"
  name                = "example"
  resource_group_name = "pike"
}
