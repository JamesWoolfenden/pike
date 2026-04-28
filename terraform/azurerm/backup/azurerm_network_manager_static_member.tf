resource "azurerm_network_manager_static_member" "pike_gen" {
  name                      = "example-nmsm"
  network_group_id          = "pike"
  target_virtual_network_id = "pike"
}
