resource "azurerm_application_load_balancer" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
}
