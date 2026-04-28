resource "azurerm_synapse_workspace_sql_aad_admin" "pike_gen" {
  synapse_workspace_id = "pike"
  login                = "AzureAD Admin"
  object_id            = "pike"
  tenant_id            = "pike"
}
