resource "azurerm_machine_learning_datastore_fileshare" "pike_gen" {
  name                 = "example-datastore"
  workspace_id         = "pike"
  storage_fileshare_id = "pike"
  account_key          = "pike"
}
