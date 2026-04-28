resource "azurerm_data_factory_linked_service_sftp" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  authentication_type = "Basic"
  host                = "http://www.bing.com"
  port                = 22
  username            = "foo"
}
