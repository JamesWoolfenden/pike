resource "azurerm_log_analytics_datasource_windows_event" "pike_gen" {
  name                = "example-lad-wpc"
  resource_group_name = "pike"
  workspace_name      = "pike"
  event_log_name      = "Application"
  event_types         = ["Error"]
}
