resource "azurerm_api_management_logger" "pike_gen" {
  name                = "example-logger"
  api_management_name = "pike"
  resource_group_name = "pike"
  resource_id         = "pike"

  application_insights {
    instrumentation_key = "pike"
  }
}
