resource "azurerm_firewall_policy" "pike_gen" {
  name                = "example-policy"
  resource_group_name = "pike"
  location            = "pike"
}
