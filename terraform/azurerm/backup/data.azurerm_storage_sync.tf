data "azurerm_storage_sync" "pike" {
  resource_group_name = "pike"
  name                = azurerm_storage_sync.pike.name
}
