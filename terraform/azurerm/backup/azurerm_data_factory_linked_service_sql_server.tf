resource "azurerm_data_factory_linked_service_sql_server" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
}
