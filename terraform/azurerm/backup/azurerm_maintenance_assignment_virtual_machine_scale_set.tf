resource "azurerm_maintenance_assignment_virtual_machine_scale_set" "pike_gen" {
  location                     = "pike"
  maintenance_configuration_id = "pike"
  virtual_machine_scale_set_id = "pike"
}
