resource "azurerm_monitor_diagnostic_setting" "pike_gen" {
  name               = "example"
  target_resource_id = "pike"
  storage_account_id = "pike"

  enabled_log {
    category = "AuditEvent"
  }

  enabled_metric {
    category = "AllMetrics"
  }
}
