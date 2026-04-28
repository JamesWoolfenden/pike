resource "azurerm_system_center_virtual_machine_manager_availability_set" "pike_gen" {
  name                                            = "example-scvmmas"
  resource_group_name                             = "pike"
  location                                        = "pike"
  custom_location_id                              = "pike"
  system_center_virtual_machine_manager_server_id = "pike"
}
