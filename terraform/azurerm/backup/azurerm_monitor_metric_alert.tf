resource "azurerm_monitor_metric_alert" "pike_gen" {
  name                = "example-metricalert"
  resource_group_name = "pike"
  scopes              = ["pike"]
  description         = "Action will be triggered when Transactions count is greater than 50."

  criteria {
    metric_namespace = "Microsoft.Storage/storageAccounts"
    metric_name      = "Transactions"
    aggregation      = "Total"
    operator         = "GreaterThan"
    threshold        = 50

    dimension {
      name     = "ApiName"
      operator = "Include"
      values   = ["*"]
    }
  }

  action {
    action_group_id = "pike"
  }
}
