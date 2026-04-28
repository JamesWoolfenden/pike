resource "azurerm_storage_container_immutability_policy" "pike_gen" {
  storage_container_resource_manager_id = "pike"
  immutability_period_in_days           = 14
  protected_append_writes_all_enabled   = false
  protected_append_writes_enabled       = true
}
