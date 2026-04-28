resource "azurerm_managed_application_definition" "pike_gen" {
  name                = "examplemanagedapplicationdefinition"
  location            = "pike"
  resource_group_name = "pike"
  lock_level          = "ReadOnly"
  package_file_uri    = "https://github.com/Azure/azure-managedapp-samples/raw/master/Managed Application Sample Packages/201-managed-storage-account/managedstorage.zip"
  display_name        = "TestManagedApplicationDefinition"
  description         = "Test Managed Application Definition"

  authorization {
    service_principal_id = "pike"
    role_definition_id   = "a094b430-dad3-424d-ae58-13f72fd72591"
  }
}
