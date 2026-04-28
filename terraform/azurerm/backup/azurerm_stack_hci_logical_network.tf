resource "azurerm_stack_hci_logical_network" "pike_gen" {
  name                = "example-hci-ln"
  resource_group_name = "pike"
  location            = "pike"
  custom_location_id  = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/cl1"
  virtual_switch_name = "ConvergedSwitch(managementcompute)"
  dns_servers         = ["10.0.0.7", "10.0.0.8"]

  subnet {
    ip_allocation_method = "Static"
    address_prefix       = "10.0.0.0/24"
    vlan_id              = 123
    route {
      address_prefix      = "0.0.0.0/0"
      next_hop_ip_address = "10.0.0.1"
    }
  }

  tags = {
    foo = "bar"
  }
}
