resource "azurerm_new_relic_tag_rule" "pike_gen" {
  monitor_id                         = "pike"
  azure_active_directory_log_enabled = true
  activity_log_enabled               = true
  metric_enabled                     = true
  subscription_log_enabled           = true

  log_tag_filter {
    name   = "key"
    action = "Include"
    value  = "value"
  }

  metric_tag_filter {
    name   = "key"
    action = "Exclude"
    value  = "value"
  }
}
