resource "azurerm_sentinel_data_connector_dynamics_365" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
}
