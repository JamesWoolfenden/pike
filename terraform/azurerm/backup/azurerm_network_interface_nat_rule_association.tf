resource "azurerm_network_interface_nat_rule_association" "pike_gen" {
  network_interface_id  = "pike"
  ip_configuration_name = "testconfiguration1"
  nat_rule_id           = "pike"
}
