resource "azurerm_data_factory_integration_runtime_azure" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  location        = "pike"
}
