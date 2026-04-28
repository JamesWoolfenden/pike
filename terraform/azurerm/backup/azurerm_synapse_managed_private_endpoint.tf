resource "azurerm_synapse_managed_private_endpoint" "pike_gen" {
  name                 = "example-endpoint"
  synapse_workspace_id = "pike"
  target_resource_id   = "pike"
  subresource_name     = "blob"

  depends_on = ["pike"]
}
