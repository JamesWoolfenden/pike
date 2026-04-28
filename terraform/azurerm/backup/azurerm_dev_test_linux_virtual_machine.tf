resource "azurerm_dev_test_linux_virtual_machine" "pike_gen" {
  name                   = "example-vm03"
  lab_name               = "pike"
  resource_group_name    = "pike"
  location               = "pike"
  size                   = "Standard_DS2"
  username               = "exampleuser99"
  ssh_key                = file("~/.ssh/id_rsa.pub")
  lab_virtual_network_id = "pike"
  lab_subnet_name        = "pike"
  storage_type           = "Premium"
  notes                  = "Some notes about this Virtual Machine."

  gallery_image_reference {
    publisher = "Canonical"
    offer     = "0001-com-ubuntu-server-jammy"
    sku       = "22_04-lts"
    version   = "latest"
  }
}
