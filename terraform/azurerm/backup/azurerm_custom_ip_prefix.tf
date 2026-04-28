resource "azurerm_custom_ip_prefix" "pike_gen" {
  name                = "example-CustomIPPrefix"
  location            = "pike"
  resource_group_name = "pike"

  cidr  = "1.2.3.4/22"
  zones = ["1", "2", "3"]

  commissioning_enabled = true

  roa_validity_end_date         = "2099-12-12"
  wan_validation_signed_message = "signed message for WAN validation"

  tags = {
    env = "test"
  }
}
