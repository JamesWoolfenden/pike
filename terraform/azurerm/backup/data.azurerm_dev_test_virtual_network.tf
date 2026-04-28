data "azurerm_dev_test_virtual_network" "pike_gen" {
  name                = "example-network"
  lab_name            = "examplelab"
  resource_group_name = "example-resource"
}
