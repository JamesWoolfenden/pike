resource "azurerm_static_web_app" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
}
