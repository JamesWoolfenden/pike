resource "azurerm_log_analytics_datasource_windows_performance_counter" "pike_gen" {
  name                = "example-lad-wpc"
  resource_group_name = "pike"
  workspace_name      = "pike"
  object_name         = "CPU"
  instance_name       = "*"
  counter_name        = "CPU"
  interval_seconds    = 10
}
