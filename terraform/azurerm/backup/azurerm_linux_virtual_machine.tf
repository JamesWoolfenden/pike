
resource "azurerm_linux_virtual_machine" "pike" {
  name                = "example-machine"
  resource_group_name = data.azurerm_resource_group.pike.name
  location            = data.azurerm_resource_group.pike.location
  size                = "Standard_F2"
  admin_username      = "adminuser"
  network_interface_ids = [
    azurerm_network_interface.pike.id,
  ]

  admin_ssh_key {
    username   = "adminuser"
    public_key = tls_private_key.pike.public_key_openssh
  }

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
}


resource "tls_private_key" "pike" {
  algorithm = "RSA"
}
