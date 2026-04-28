resource "azurerm_log_analytics_linked_storage_account" "pike_gen" {
  data_source_type    = "CustomLogs"
  resource_group_name = "pike"
  workspace_id        = "pike"
  storage_account_ids = ["pike"]
}
