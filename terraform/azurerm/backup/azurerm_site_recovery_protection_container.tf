resource "azurerm_site_recovery_protection_container" "pike_gen" {
  name                 = "protection-container"
  resource_group_name  = "pike"
  recovery_vault_name  = "pike"
  recovery_fabric_name = "pike"
}
