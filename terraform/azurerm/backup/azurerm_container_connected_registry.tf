resource "azurerm_container_connected_registry" "pike_gen" {
  name                  = "examplecr"
  container_registry_id = "pike"
  sync_token_id         = "pike"
}
