resource "azurerm_container_app_environment" "pike_gen" {
  name                       = "my-environment"
  location                   = "pike"
  resource_group_name        = "pike"
  logs_destination           = "log-analytics"
  log_analytics_workspace_id = "pike"
}
