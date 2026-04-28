resource "azurerm_storage_sync_server_endpoint" "pike_gen" {
  name                  = "example-storage-sync-server-endpoint"
  storage_sync_group_id = "pike"
  registered_server_id  = "pike"
}
