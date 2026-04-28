resource "azurerm_data_factory_dataset_postgresql" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"
}
