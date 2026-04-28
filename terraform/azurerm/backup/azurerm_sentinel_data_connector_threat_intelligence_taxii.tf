resource "azurerm_sentinel_data_connector_threat_intelligence_taxii" "pike_gen" {
  name                       = "example"
  log_analytics_workspace_id = "pike"
  display_name               = "example"
  api_root_url               = "https://foo/taxii2/api2/"
  collection_id              = "someid"
}
