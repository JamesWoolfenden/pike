resource "azurerm_machine_learning_compute_cluster" "pike_gen" {
  name                          = "example"
  location                      = "pike"
  vm_priority                   = "LowPriority"
  vm_size                       = "Standard_DS2_v2"
  machine_learning_workspace_id = "pike"
  subnet_resource_id            = "pike"

  scale_settings {
    min_node_count                       = 0
    max_node_count                       = 1
    scale_down_nodes_after_idle_duration = "PT30S" # 30 seconds
  }

  identity {
    type = "SystemAssigned"
  }
}
