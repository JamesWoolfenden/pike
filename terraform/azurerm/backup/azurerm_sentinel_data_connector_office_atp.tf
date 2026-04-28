resource "azurerm_sentinel_data_connector_office_atp" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
