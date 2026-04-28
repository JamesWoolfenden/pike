resource "azurerm_data_factory_linked_service_kusto" "pike_gen" {
  name                 = "example"
  data_factory_id      = "pike"
  kusto_endpoint       = "pike"
  kusto_database_name  = "pike"
  use_managed_identity = true
}
