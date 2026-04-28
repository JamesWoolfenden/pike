resource "azurerm_monitor_scheduled_query_rules_log" "pike_gen" {
  name                = "example"
  location            = "pike"
  resource_group_name = "pike"

  criteria {
    metric_name = "Average_% Idle Time"
    dimension {
      name     = "Computer"
      operator = "Include"
      values   = ["targetVM"]
    }
  }
  data_source_id = "pike"
  description    = "Scheduled query rule LogToMetric example"
  enabled        = true
  tags = {
    foo = "bar"
  }
}
