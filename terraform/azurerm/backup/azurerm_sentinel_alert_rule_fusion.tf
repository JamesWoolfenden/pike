resource "azurerm_sentinel_alert_rule_fusion" "pike_gen" {
  log_analytics_workspace_id = "pike"
  alert_rule_template_guid   = "f71aba3d-28fb-450b-b192-4e76a83015c8"
}
