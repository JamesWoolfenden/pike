resource "azurerm_log_analytics_linked_service" "pike_gen" {
  resource_group_name = "pike"
  workspace_id        = "pike"
  read_access_id      = "pike"
}
