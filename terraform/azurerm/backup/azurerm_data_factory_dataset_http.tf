resource "azurerm_data_factory_dataset_http" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"

  relative_url   = "http://www.bing.com"
  request_body   = "foo=bar"
  request_method = "POST"

}
