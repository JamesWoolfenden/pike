resource "azurerm_data_factory_linked_service_odbc" "pike_gen" {
  name            = "anonymous"
  data_factory_id = "pike"
}
