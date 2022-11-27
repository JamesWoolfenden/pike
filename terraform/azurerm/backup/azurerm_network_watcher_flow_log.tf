resource "azurerm_network_watcher_flow_log" "pike" {
  network_watcher_name = "NetworkWatcher_uksouth"
  resource_group_name  = "NetworkWatcherRG"
  name                 = "example-log"

  network_security_group_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/NetworkWatcherRG/providers/Microsoft.Network/networkSecurityGroups/pike"
  storage_account_id        = azurerm_storage_account.pike.id
  enabled                   = true

  retention_policy {
    enabled = true
    days    = 7
  }

  traffic_analytics {
    enabled               = true
    workspace_id          = azurerm_log_analytics_workspace.pike.workspace_id
    workspace_region      = azurerm_log_analytics_workspace.pike.location
    workspace_resource_id = azurerm_log_analytics_workspace.pike.id
    interval_in_minutes   = 10
  }
}
