resource "azurerm_sentinel_data_connector_office_irm" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
