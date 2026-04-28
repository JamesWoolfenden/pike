resource "azurerm_data_factory_dataset_delimited_text" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"

  http_server_location {
    relative_url = "http://www.bing.com"
    path         = "foo/bar/"
    filename     = "fizz.txt"
  }

  column_delimiter    = ","
  row_delimiter       = "NEW"
  encoding            = "UTF-8"
  quote_character     = "x"
  escape_character    = "f"
  first_row_as_header = true
  null_value          = "NULL"
}
