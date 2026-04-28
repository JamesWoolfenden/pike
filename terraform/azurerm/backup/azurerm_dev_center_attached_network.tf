resource "azurerm_dev_center_attached_network" "pike_gen" {
  name                  = "example-dcet"
  dev_center_id         = "pike"
  network_connection_id = "pike"
}
