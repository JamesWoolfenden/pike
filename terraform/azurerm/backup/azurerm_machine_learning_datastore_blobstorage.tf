resource "azurerm_machine_learning_datastore_blobstorage" "pike_gen" {
  name                 = "example-datastore"
  workspace_id         = "pike"
  storage_container_id = "pike"
  account_key          = "pike"
}
