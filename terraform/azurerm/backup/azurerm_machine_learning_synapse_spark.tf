resource "azurerm_machine_learning_synapse_spark" "pike_gen" {
  name                          = "example"
  machine_learning_workspace_id = "pike"
  location                      = "pike"
  synapse_spark_pool_id         = "pike"

  identity {
    type = "SystemAssigned"
  }
}
