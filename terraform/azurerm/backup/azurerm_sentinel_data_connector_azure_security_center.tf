resource "azurerm_sentinel_data_connector_azure_security_center" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
