resource "azurerm_monitor_smart_detector_alert_rule" "pike_gen" {
  name                = "example-smart-detector-alert-rule"
  resource_group_name = "pike"
  severity            = "Sev0"
  scope_resource_ids  = ["pike"]
  frequency           = "PT1M"
  detector_type       = "FailureAnomaliesDetector"

  action_group {
    ids = ["pike"]
  }
}
