resource "azurerm_synapse_workspace_extended_auditing_policy" "pike_gen" {
  synapse_workspace_id = "pike"
  storage_endpoint     = "pike"
  retention_in_days    = 6
}
