resource "azurerm_sentinel_automation_rule" "pike_gen" {
  name                       = "56094f72-ac3f-40e7-a0c0-47bd95f70336"
  log_analytics_workspace_id = "pike"
  display_name               = "automation_rule1"
  order                      = 1
  action_incident {
    order  = 1
    status = "Active"
  }
}
