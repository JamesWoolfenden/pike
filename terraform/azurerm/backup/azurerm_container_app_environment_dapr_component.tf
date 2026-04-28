resource "azurerm_container_app_environment_dapr_component" "pike_gen" {
  name                         = "example-component"
  container_app_environment_id = "pike"
  component_type               = "state.azure.blobstorage"
  version                      = "v1"
}
