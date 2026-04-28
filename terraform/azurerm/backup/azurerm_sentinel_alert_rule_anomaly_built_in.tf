resource "azurerm_sentinel_alert_rule_anomaly_built_in" "pike_gen" {
  display_name               = "Potential data staging"
  log_analytics_workspace_id = "pike"
  mode                       = "Production"
  enabled                    = false
}
