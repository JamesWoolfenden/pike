resource "azurerm_application_insights_standard_web_test" "pike_gen" {
  name                    = "example-test"
  resource_group_name     = "pike"
  location                = "West Europe"
  application_insights_id = "pike"
  geo_locations           = ["example"]

  request {
    url = "http://www.example.com"
  }
}
