data "azurerm_network_security_perimeter_profile" "pike_gen" {
  name                          = "existing"
  network_security_perimeter_id = "pike"
}
