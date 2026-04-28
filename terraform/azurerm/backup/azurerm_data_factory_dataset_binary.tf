resource "azurerm_data_factory_dataset_binary" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"

  sftp_server_location {
    path     = "/test/"
    filename = "**"
  }
}
