resource "azurerm_sentinel_data_connector_iot" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
