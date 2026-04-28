resource "azurerm_function_app_connection" "pike_gen" {
  name               = "example-serviceconnector"
  function_app_id    = "pike"
  target_resource_id = "pike"
  authentication {
    type = "systemAssignedIdentity"
  }
}
