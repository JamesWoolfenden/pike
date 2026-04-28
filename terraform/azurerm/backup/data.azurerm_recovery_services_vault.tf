data "azurerm_recovery_services_vault" "pike_gen" {
  name                = "tfex-recovery_vault"
  resource_group_name = "tfex-resource_group"
}
