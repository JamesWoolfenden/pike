resource "azurerm_subnet_network_security_group_association" "pike_gen" {
  subnet_id                 = "pike"
  network_security_group_id = "pike"
}
