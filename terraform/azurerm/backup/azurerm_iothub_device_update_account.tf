resource "azurerm_iothub_device_update_account" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"

  identity {
    type = "SystemAssigned"
  }

  tags = {
    key = "value"
  }
}
