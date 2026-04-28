resource "azurerm_recovery_services_vault" "pike_gen" {
  name                = "example-recovery-vault"
  location            = "pike"
  resource_group_name = "pike"
  sku                 = "Standard"
}
