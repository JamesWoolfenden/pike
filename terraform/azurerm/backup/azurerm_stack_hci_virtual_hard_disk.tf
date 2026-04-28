resource "azurerm_stack_hci_virtual_hard_disk" "pike_gen" {
  name                = "example-vhd"
  resource_group_name = "pike"
  location            = "pike"
  custom_location_id  = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/cl1"
  disk_size_in_gb     = 2
  storage_path_id     = "pike"
  tags = {
    foo = "bar"
  }
}
