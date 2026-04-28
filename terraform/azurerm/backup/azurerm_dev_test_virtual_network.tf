resource "azurerm_dev_test_virtual_network" "pike_gen" {
  name                = "example-network"
  lab_name            = "pike"
  resource_group_name = "pike"

  subnet {
    use_public_ip_address           = "Allow"
    use_in_virtual_machine_creation = "Allow"
  }
}
