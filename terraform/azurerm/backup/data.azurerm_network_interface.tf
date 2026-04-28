data "azurerm_network_interface" "pike_gen" {
  name                = "acctest-nic"
  resource_group_name = "networking"
}
