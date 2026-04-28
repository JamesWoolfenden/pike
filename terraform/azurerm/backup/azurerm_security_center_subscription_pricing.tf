resource "azurerm_security_center_subscription_pricing" "pike_gen" {
  tier          = "Standard"
  resource_type = "VirtualMachines"
}
