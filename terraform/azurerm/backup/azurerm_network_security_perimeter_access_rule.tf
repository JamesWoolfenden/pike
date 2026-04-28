resource "azurerm_network_security_perimeter_access_rule" "pike_gen" {
  name                                  = "example"
  network_security_perimeter_profile_id = "pike"
  direction                             = "Inbound"

  address_prefixes = [
    "8.8.8.8/32"
  ]
}
