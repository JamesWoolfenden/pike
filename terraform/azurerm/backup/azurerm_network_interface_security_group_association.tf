resource "azurerm_network_interface_security_group_association" "pike_gen" {
  network_interface_id      = "pike"
  network_security_group_id = "pike"
}
