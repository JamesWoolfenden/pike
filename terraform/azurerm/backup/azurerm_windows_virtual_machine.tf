resource "azurerm_windows_virtual_machine" "example" {
  name                = "example-win"
  resource_group_name = data.azurerm_resource_group.pike.name
  location            = data.azurerm_resource_group.pike.location
  size                = "Standard_F2"
  admin_username      = "adminuser"
  admin_password      = "P@$$w0rd1234!"
  network_interface_ids = [
    azurerm_network_interface.win.id,
  ]

  os_disk {
    caching              = "ReadWrite"
    storage_account_type = "Standard_LRS"
  }

  source_image_reference {
    publisher = "MicrosoftWindowsServer"
    offer     = "WindowsServer"
    sku       = "2016-Datacenter"
    version   = "latest"
  }
}

resource "azurerm_network_interface" "win" {
  name                = "example-win"
  location            = data.azurerm_resource_group.pike.location
  resource_group_name = data.azurerm_resource_group.pike.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = data.azurerm_subnet.pike.id
    private_ip_address_allocation = "Dynamic"
  }
}
