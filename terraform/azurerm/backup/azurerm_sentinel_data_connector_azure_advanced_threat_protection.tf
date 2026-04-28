resource "azurerm_sentinel_data_connector_azure_advanced_threat_protection" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
