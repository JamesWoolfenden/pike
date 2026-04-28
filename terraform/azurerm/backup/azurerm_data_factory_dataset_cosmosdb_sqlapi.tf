resource "azurerm_data_factory_dataset_cosmosdb_sqlapi" "pike_gen" {
  name                = "example"
  data_factory_id     = "pike"
  linked_service_name = "pike"

  collection_name = "bar"
}
