resource "azurerm_subnet_service_endpoint_storage_policy" "pike_gen" {
  name                = "example-policy"
  resource_group_name = "pike"
  location            = "pike"
  definition {
    name        = "name1"
    description = "definition1"
    service     = "Microsoft.Storage"
    service_resources = [
      azurerm_resource_group.example.id,
      azurerm_storage_account.example.id
    ]
  }
  definition {
    name        = "name2"
    description = "definition2"
    service     = "Global"
    service_resources = [
      "/services/Azure",
      "/services/Azure/Batch",
      "/services/Azure/Databricks",
      "/services/Azure/DataFactory",
      "/services/Azure/MachineLearning",
      "/services/Azure/ManagedInstance",
      "/services/Azure/WebPI",
    ]
  }
}
