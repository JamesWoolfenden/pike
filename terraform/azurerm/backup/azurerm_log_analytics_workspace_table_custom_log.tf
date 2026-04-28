resource "azurerm_log_analytics_workspace_table_custom_log" "pike_gen" {
  name         = "example_CL"
  workspace_id = "pike"

  column {
    name = "TimeGenerated"
    type = "datetime"
  }
}
