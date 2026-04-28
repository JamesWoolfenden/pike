resource "azurerm_palo_alto_local_rulestack" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
}
