resource "azurerm_machine_learning_datastore_datalake_gen2" "pike_gen" {
  name                 = "example-datastore"
  workspace_id         = "pike"
  storage_container_id = "pike"
}
