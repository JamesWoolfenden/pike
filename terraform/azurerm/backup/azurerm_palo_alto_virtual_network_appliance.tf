resource "azurerm_palo_alto_virtual_network_appliance" "pike_gen" {
  name           = "example-appliance"
  virtual_hub_id = "pike"
}
