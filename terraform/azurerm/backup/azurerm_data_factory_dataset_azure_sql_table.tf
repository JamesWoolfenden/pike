resource "azurerm_data_factory_dataset_azure_sql_table" "pike_gen" {
  name              = "example"
  data_factory_id   = "pike"
  linked_service_id = "pike"
}
