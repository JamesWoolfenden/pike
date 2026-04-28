resource "azurerm_virtual_machine_automanage_configuration_assignment" "pike_gen" {
  virtual_machine_id = "pike"
  configuration_id   = "pike"
}
