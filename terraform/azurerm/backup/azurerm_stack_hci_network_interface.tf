resource "azurerm_stack_hci_network_interface" "pike_gen" {
  name                = "example-ni"
  resource_group_name = "pike"
  location            = "pike"
  custom_location_id  = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/cl1"
  dns_servers         = ["10.0.0.8"]

  ip_configuration {
    private_ip_address = "10.0.0.2"
    subnet_id          = "pike"
  }

  tags = {
    foo = "bar"
  }

  lifecycle {
    ignore_changes = [mac_address]
  }
}
