resource "azurerm_log_analytics_data_export_rule" "pike_gen" {
  name                    = "dataExport1"
  resource_group_name     = "pike"
  workspace_resource_id   = "pike"
  destination_resource_id = "pike"
  table_names             = ["Heartbeat"]
  enabled                 = true
}
