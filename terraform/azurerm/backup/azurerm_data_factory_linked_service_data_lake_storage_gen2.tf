resource "azurerm_data_factory_linked_service_data_lake_storage_gen2" "pike_gen" {
  name                  = "example"
  data_factory_id       = "pike"
  service_principal_id  = "pike"
  service_principal_key = "exampleKey"
  tenant                = "11111111-1111-1111-1111-111111111111"
  url                   = "https://datalakestoragegen2"
}
