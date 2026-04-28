resource "azurerm_application_insights_api_key" "pike_gen" {
  name                    = "tf-test-appinsights-read-telemetry-api-key"
  application_insights_id = "pike"
  read_permissions        = ["aggregate", "api", "draft", "extendqueries", "search"]
}
