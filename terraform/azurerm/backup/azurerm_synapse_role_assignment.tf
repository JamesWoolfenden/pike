resource "azurerm_synapse_role_assignment" "pike_gen" {
  synapse_workspace_id = "pike"
  role_name            = "Synapse SQL Administrator"
  principal_id         = "pike"

  depends_on = ["pike"]
}
