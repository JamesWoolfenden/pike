resource "azurerm_maintenance_assignment_virtual_machine" "pike_gen" {
  location                     = "pike"
  maintenance_configuration_id = "pike"
  virtual_machine_id           = "pike"
}
