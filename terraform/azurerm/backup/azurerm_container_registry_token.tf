resource "azurerm_container_registry_token" "pike_gen" {
  name                    = "exampletoken"
  container_registry_name = "pike"
  resource_group_name     = "pike"
  scope_map_id            = "pike"
}
