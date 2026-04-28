resource "azurerm_data_factory_credential_service_principal" "pike_gen" {
  name                 = "example"
  description          = "example description"
  data_factory_id      = "pike"
  tenant_id            = "pike"
  service_principal_id = "pike"
  service_principal_key {
    linked_service_name = "pike"
    secret_name         = "pike"
    secret_version      = "pike"
  }
  annotations = ["1", "2"]
}
