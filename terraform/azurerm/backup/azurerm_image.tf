resource "azurerm_image" "pike" {
  name                      = "pike-image"
  location                  = "uksouth"
  resource_group_name       = "pike"
  source_virtual_machine_id = azurerm_virtual_machine.pike.id

  tags = {
    pike = "permissions"
  }
}
