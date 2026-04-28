resource "azurerm_stack_hci_storage_path" "pike_gen" {
  name                = "example-sp"
  resource_group_name = "pike"
  location            = "pike"
  custom_location_id  = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/cl1"
  path                = "C:\\ClusterStorage\\UserStorage_2\\sp-example"
  tags = {
    foo = "bar"
  }
}
