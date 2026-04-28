resource "azurerm_network_ddos_protection_plan" "pike_gen" {
  name                = "example-protection-plan"
  location            = "pike"
  resource_group_name = "pike"
}
