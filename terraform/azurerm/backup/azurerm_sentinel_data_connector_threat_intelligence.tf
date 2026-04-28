resource "azurerm_sentinel_data_connector_threat_intelligence" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
