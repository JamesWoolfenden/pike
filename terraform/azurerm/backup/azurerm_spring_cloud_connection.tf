resource "azurerm_spring_cloud_connection" "pike_gen" {
  name               = "example-serviceconnector"
  spring_cloud_id    = "pike"
  target_resource_id = "pike"
  authentication {
    type = "systemAssignedIdentity"
  }
}
