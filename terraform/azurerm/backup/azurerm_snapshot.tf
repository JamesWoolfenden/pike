resource "azurerm_snapshot" "pike" {
  name                = "pike-snapshot"
  location            = "uksouth"
  resource_group_name = "pike"
  create_option       = "Copy"
  source_uri          = azurerm_managed_disk.pike.id

  tags = {
    pike = "permissions"
  }
}
