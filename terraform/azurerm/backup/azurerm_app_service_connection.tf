resource "azurerm_app_service_connection" "pike_gen" {
  name               = "example-serviceconnector"
  app_service_id     = "pike"
  target_resource_id = "pike"
  authentication {
    type = "systemAssignedIdentity"
  }
}
