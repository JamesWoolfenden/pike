resource "azurerm_sentinel_data_connector_microsoft_threat_intelligence" "pike_gen" {
  name                                         = "example-dc-msti"
  log_analytics_workspace_id                   = "pike"
  microsoft_emerging_threat_feed_lookback_date = "1970-01-01T00:00:00Z"
}
