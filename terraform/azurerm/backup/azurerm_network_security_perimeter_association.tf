resource "azurerm_network_security_perimeter_association" "pike_gen" {
  name        = "example"
  access_mode = "Enforced"

  network_security_perimeter_profile_id = "pike"
  resource_id                           = "pike"
}
