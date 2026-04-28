resource "azurerm_application_insights_analytics_item" "pike_gen" {
  name                    = "testquery"
  application_insights_id = "pike"
  content                 = "requests //simple example query"
  scope                   = "shared"
  type                    = "query"
}
