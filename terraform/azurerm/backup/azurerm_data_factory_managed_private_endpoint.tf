resource "azurerm_data_factory_managed_private_endpoint" "pike_gen" {
  name               = "example"
  data_factory_id    = "pike"
  target_resource_id = "pike"
  subresource_name   = "blob"
}
