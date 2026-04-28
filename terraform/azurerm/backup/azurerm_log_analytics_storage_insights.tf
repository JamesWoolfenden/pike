resource "azurerm_log_analytics_storage_insights" "pike_gen" {
  name                = "example-storageinsightconfig"
  resource_group_name = "pike"
  workspace_id        = "pike"

  storage_account_id  = "pike"
  storage_account_key = "pike"
}
