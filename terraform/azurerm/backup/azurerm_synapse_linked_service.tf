resource "azurerm_synapse_linked_service" "pike_gen" {
  name                 = "example"
  synapse_workspace_id = "pike"
  type                 = "AzureBlobStorage"
  type_properties_json = <<JSON
{
  "connectionString": "${azurerm_storage_account.example.primary_connection_string}"
}
JSON
  integration_runtime {
    name = "pike"
  }

  depends_on = [
    azurerm_synapse_firewall_rule.example,
  ]
}
