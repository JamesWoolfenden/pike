resource "azurerm_sentinel_metadata" "pike_gen" {
  name         = "exampl"
  workspace_id = "pike"
  content_id   = "pike"
  kind         = "AnalyticsRule"
  parent_id    = "pike"
}
