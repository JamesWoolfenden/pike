resource "azurerm_network_interface_application_security_group_association" "pike_gen" {
  network_interface_id          = "pike"
  application_security_group_id = "pike"
}
