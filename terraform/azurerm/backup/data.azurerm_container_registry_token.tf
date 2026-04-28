data "azurerm_container_registry_token" "pike_gen" {
  name                    = "exampletoken"
  resource_group_name     = "example-resource-group"
  container_registry_name = "example-registry"
}
