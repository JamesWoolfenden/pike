resource "azurerm_sentinel_alert_rule_ms_security_incident" "pike_gen" {
  name                       = "example-ms-security-incident-alert-rule"
  log_analytics_workspace_id = "pike"
  product_filter             = "Microsoft Cloud App Security"
  display_name               = "example rule"
  severity_filter            = ["High"]
}
