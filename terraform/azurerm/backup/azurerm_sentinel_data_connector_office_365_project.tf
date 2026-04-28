resource "azurerm_sentinel_data_connector_office_365_project" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
