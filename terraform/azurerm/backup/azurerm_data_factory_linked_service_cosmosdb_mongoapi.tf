resource "azurerm_data_factory_linked_service_cosmosdb_mongoapi" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  database        = "foo"

}
