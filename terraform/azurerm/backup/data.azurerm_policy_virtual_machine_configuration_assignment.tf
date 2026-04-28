data "azurerm_policy_virtual_machine_configuration_assignment" "pike_gen" {
  name                 = "AzureWindowsBaseline"
  resource_group_name  = "example-RG"
  virtual_machine_name = "example-vm"
}
