resource "azurerm_sentinel_data_connector_azure_active_directory" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
