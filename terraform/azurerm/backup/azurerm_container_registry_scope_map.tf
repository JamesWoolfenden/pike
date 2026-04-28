resource "azurerm_container_registry_scope_map" "pike_gen" {
  name                    = "example-scope-map"
  container_registry_name = "pike"
  resource_group_name     = "pike"
  actions = [
    "repositories/repo1/content/read",
    "repositories/repo1/content/write"
  ]
}
