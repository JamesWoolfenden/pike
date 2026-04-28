resource "azurerm_storage_object_replication" "pike_gen" {
  source_storage_account_id      = "pike"
  destination_storage_account_id = "pike"
  rules {
    source_container_name      = "pike"
    destination_container_name = "pike"
  }
}
