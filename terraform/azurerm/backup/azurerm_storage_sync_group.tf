resource "azurerm_storage_sync_group" "pike" {
  name            = "pike"
  storage_sync_id = azurerm_storage_sync.pike.id
}
