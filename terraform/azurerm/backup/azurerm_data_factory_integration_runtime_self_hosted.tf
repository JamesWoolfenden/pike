resource "azurerm_data_factory_integration_runtime_self_hosted" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
}
