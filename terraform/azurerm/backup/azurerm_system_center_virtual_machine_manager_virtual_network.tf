resource "azurerm_system_center_virtual_machine_manager_virtual_network" "pike_gen" {
  name                                                           = "example-scvmmvnet"
  resource_group_name                                            = "pike"
  location                                                       = "pike"
  custom_location_id                                             = "pike"
  system_center_virtual_machine_manager_server_inventory_item_id = "pike"
}
