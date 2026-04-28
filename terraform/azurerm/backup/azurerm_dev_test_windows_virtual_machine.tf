resource "azurerm_dev_test_windows_virtual_machine" "pike_gen" {
  name                   = "example-vm03"
  lab_name               = "pike"
  resource_group_name    = "pike"
  location               = "pike"
  size                   = "Standard_DS2"
  username               = "exampleuser99"
  lab_virtual_network_id = "pike"
  lab_subnet_name        = "pike"
  storage_type           = "Premium"
  notes                  = "Some notes about this Virtual Machine."

  gallery_image_reference {
    offer     = "WindowsServer"
    publisher = "MicrosoftWindowsServer"
    sku       = "2019-Datacenter"
    version   = "latest"
  }
}
