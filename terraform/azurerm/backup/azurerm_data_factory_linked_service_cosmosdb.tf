resource "azurerm_data_factory_linked_service_cosmosdb" "pike_gen" {
  name             = "example"
  data_factory_id  = "pike"
  account_endpoint = "pike"
  account_key      = "pike"
  database         = "foo"

}
