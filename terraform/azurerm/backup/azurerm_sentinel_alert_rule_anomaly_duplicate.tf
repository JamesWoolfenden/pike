resource "azurerm_sentinel_alert_rule_anomaly_duplicate" "pike_gen" {
  display_name               = "example duplicated UEBA Anomalous Sign In"
  log_analytics_workspace_id = "pike"
  built_in_rule_id           = "pike"
  enabled                    = true
  mode                       = "Flighting"

  threshold_observation {
    name  = "Anomaly score threshold"
    value = "0.6"
  }
}
