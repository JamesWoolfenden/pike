data "azurerm_sentinel_alert_rule_template" "pike" {
  log_analytics_workspace_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.OperationalInsights/workspaces/pike"
  name                       = "pike"
}
