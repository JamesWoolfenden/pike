resource "azurerm_network_manager_scope_connection" "pike_gen" {
  name               = "example-nsc"
  network_manager_id = "pike"
  tenant_id          = "pike"
  target_scope_id    = "pike"
  description        = "example"
}
