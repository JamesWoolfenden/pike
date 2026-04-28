resource "azurerm_stack_hci_marketplace_gallery_image" "pike_gen" {
  name                = "example-mgi"
  resource_group_name = "pike"
  location            = "pike"
  custom_location_id  = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/cl1"
  hyperv_generation   = "V2"
  os_type             = "Windows"
  version             = "20348.2655.240905"
  identifier {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2022-datacenter-azure-edition-core"
  }
  tags = {
    foo = "bar"
    env = "example"
  }
}
