resource "azurerm_recovery_services_vault_resource_guard_association" "pike_gen" {
  vault_id          = "pike"
  resource_guard_id = "pike"
}
