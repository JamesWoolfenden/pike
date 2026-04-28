resource "azurerm_managed_application" "pike_gen" {
  name                        = "example-managedapplication"
  location                    = "pike"
  resource_group_name         = "pike"
  kind                        = "ServiceCatalog"
  managed_resource_group_name = "infrastructureGroup"
  application_definition_id   = "pike"

  parameter_values = jsonencode({
    location = {
      value = "pike"
    },
    storageAccountNamePrefix = {
      value = "storeNamePrefix"
    },
    storageAccountType = {
      value = "Standard_LRS"
    }
  })
}
