resource "azurerm_system_center_virtual_machine_manager_server" "pike_gen" {
  name                = "example-scvmmms"
  resource_group_name = "pike"
  location            = "pike"
  custom_location_id  = "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.ExtendedLocation/customLocations/customLocation1"
  fqdn                = "example.labtest"
  username            = "testUser"
}
