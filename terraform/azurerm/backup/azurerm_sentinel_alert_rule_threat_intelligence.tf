resource "azurerm_sentinel_alert_rule_threat_intelligence" "pike_gen" {
  name                       = "example-rule"
  log_analytics_workspace_id = "pike"
  alert_rule_template_guid   = "pike"
}
