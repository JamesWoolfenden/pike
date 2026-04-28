resource "azurerm_data_factory_dataset_json" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"

  http_server_location {
    relative_url = "/fizz/buzz/"
    path         = "foo/bar/"
    filename     = "foo.txt"
  }

  encoding = "UTF-8"

}
