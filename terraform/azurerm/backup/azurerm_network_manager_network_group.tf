resource "azurerm_network_manager_network_group" "pike_gen" {
  name               = "example-group"
  network_manager_id = "pike"
}
