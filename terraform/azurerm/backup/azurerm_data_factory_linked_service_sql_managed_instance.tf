resource "azurerm_data_factory_linked_service_sql_managed_instance" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
}
