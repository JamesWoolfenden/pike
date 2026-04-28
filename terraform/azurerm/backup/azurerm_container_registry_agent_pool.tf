resource "azurerm_container_registry_agent_pool" "pike_gen" {
  name                    = "example"
  resource_group_name     = "pike"
  location                = "pike"
  container_registry_name = "pike"
}
