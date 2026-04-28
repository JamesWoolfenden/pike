resource "azurerm_synapse_integration_runtime_self_hosted" "pike_gen" {
  name                 = "example"
  synapse_workspace_id = "pike"
}
