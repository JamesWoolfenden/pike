resource "azurerm_synapse_integration_runtime_azure" "pike_gen" {
  name                 = "example"
  synapse_workspace_id = "pike"
  location             = "pike"
}
