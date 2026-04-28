resource "azurerm_virtual_hub_connection" "pike_gen" {
  name                      = "example-vhub"
  virtual_hub_id            = "pike"
  remote_virtual_network_id = "pike"
}
