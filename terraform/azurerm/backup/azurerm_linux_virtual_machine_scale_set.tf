
resource "azurerm_linux_virtual_machine_scale_set" "pike" {
  name                = "example-vmss"
  resource_group_name = "pike"
  location            = "uksouth"
  sku                 = "Standard_F2"
  instances           = 1
  admin_username      = "adminuser"

  admin_ssh_key {
    username   = "adminuser"
    public_key = tls_private_key.pike.public_key_openssh
  }

  source_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }

  os_disk {
    storage_account_type = "Standard_LRS"
    caching              = "ReadWrite"
  }

  network_interface {
    name    = "example"
    primary = true

    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = azurerm_subnet.pike.id
    }
  }
}


resource "tls_private_key" "pike" {
  algorithm = "RSA"
}
