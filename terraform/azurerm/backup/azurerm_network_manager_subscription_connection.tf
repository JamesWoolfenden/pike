resource "azurerm_network_manager_subscription_connection" "pike_gen" {
  name               = "example-nsnmc"
  subscription_id    = "pike"
  network_manager_id = "pike"
  description        = "example"
}
