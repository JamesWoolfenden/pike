resource "azurerm_synapse_private_link_hub" "pike_gen" {
  name                = "example"
  resource_group_name = "example-rg"
  location            = "West Europe"
}
