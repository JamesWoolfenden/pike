resource "azurerm_sentinel_log_analytics_workspace_onboarding" "pike_gen" {
  workspace_id                 = "pike"
  customer_managed_key_enabled = false
}
