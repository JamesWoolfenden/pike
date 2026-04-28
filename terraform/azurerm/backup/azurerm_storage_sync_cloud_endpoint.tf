resource "azurerm_storage_sync_cloud_endpoint" "pike_gen" {
  name                  = "example-ss-ce"
  storage_sync_group_id = "pike"
  file_share_name       = "pike"
  storage_account_id    = "pike"
}
