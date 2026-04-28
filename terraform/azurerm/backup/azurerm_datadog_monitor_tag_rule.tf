resource "azurerm_datadog_monitor_tag_rule" "pike_gen" {
  datadog_monitor_id = "pike"
  log {
    subscription_log_enabled = true
  }
  metric {
    filter {
      name   = "Test"
      value  = "Logs"
      action = "Include"
    }
  }
}
