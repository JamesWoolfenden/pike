data "azurerm_storage_sync_group" "pike" {
  storage_sync_id = azurerm_storage_sync_group.pike.storage_sync_id
  name            = "pike"
}
