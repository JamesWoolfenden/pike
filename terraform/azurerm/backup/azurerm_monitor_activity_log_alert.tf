resource "azurerm_monitor_activity_log_alert" "pike_gen" {
  name                = "example-activitylogalert"
  resource_group_name = "pike"
  location            = "pike"
  scopes              = ["pike"]
  description         = "This alert will monitor a specific storage account updates."

  criteria {
    resource_id    = "pike"
    operation_name = "Microsoft.Storage/storageAccounts/write"
    category       = "Recommendation"
  }

  action {
    action_group_id = "pike"

    webhook_properties = {
      from = "terraform"
    }
  }
}
