data "azurerm_container_registry_scope_map" "pike_gen" {
  name                    = "example-scope-map"
  resource_group_name     = "example-resource-group"
  container_registry_name = "example-registry"
}
