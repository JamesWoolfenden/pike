resource "azurerm_sentinel_threat_intelligence_indicator" "pike_gen" {
  workspace_id      = "pike"
  pattern_type      = "domain-name"
  pattern           = "http://example.com"
  source            = "Microsoft Sentinel"
  validate_from_utc = "2022-12-14T16:00:00Z"
  display_name      = "example-indicator"
}
