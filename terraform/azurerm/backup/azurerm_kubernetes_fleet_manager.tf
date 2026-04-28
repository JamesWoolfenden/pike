resource "azurerm_kubernetes_fleet_manager" "pike_gen" {
  location            = "pike"
  name                = "example"
  resource_group_name = "pike"
}
