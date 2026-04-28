resource "azurerm_data_factory_linked_service_web" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  authentication_type = "Anonymous"
  url                 = "http://www.bing.com"
}
