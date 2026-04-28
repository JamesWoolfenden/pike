resource "azurerm_data_factory_integration_runtime_azure_ssis" "pike_gen" {
  name            = "example"
  data_factory_id = "pike"
  location        = "pike"

  node_size = "Standard_D8_v3"
}
