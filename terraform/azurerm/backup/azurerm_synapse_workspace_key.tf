resource "azurerm_synapse_workspace_key" "pike_gen" {
  customer_managed_key_versionless_id = "pike"
  synapse_workspace_id                = "pike"
  active                              = true
  customer_managed_key_name           = "enckey"
  depends_on                          = ["pike"]
}
