resource "azurerm_virtual_machine" "main" {
  name                  = "unknown-vm"
  location              = data.azurerm_resource_group.pike.location
  resource_group_name   = data.azurerm_resource_group.pike.name
  network_interface_ids = [azurerm_network_interface.vague.id]
  vm_size               = "Standard_DS1_v2"

  # Uncomment this line to delete the OS disk automatically when deleting the VM
  # delete_os_disk_on_termination = true

  # Uncomment this line to delete the data disks automatically when deleting the VM
  # delete_data_disks_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
  storage_os_disk {
    name              = "myosdisk1"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }
  os_profile {
    computer_name  = "hostname"
    admin_username = "testadmin"
    admin_password = "Password1234!"
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
  tags = {
    pike = "permissions"
  }
}

resource "azurerm_network_interface" "vague" {
  name                = "example-vague"
  location            = data.azurerm_resource_group.pike.location
  resource_group_name = data.azurerm_resource_group.pike.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = data.azurerm_subnet.pike.id
    private_ip_address_allocation = "Dynamic"
  }
}
