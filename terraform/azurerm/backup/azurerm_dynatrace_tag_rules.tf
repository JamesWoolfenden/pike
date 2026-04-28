resource "azurerm_dynatrace_tag_rules" "pike_gen" {
  name       = "default"
  monitor_id = "pike"

  log_rule {
    filtering_tag {
      name   = "Environment"
      value  = "Prod"
      action = "Include"
    }
    send_azure_active_directory_logs_enabled = true
    send_activity_logs_enabled               = true
    send_subscription_logs_enabled           = true
  }

  metric_rule {
    filtering_tag {
      name   = "Environment"
      value  = "Prod"
      action = "Include"
    }
  }
}
