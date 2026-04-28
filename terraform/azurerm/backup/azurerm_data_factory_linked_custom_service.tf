resource "azurerm_data_factory_linked_custom_service" "pike_gen" {
  name                 = "example"
  data_factory_id      = "pike"
  type                 = "AzureBlobStorage"
  description          = "test description"
  type_properties_json = <<JSON
{
  "connectionString":"${azurerm_storage_account.example.primary_connection_string}"
}
JSON

  parameters = {
    "foo" : "bar"
    "Env" : "Test"
  }

  annotations = [
    "test1",
    "test2",
    "test3"
  ]
}
