resource "azurerm_kubernetes_cluster_node_pool" "pike_gen" {
  name                  = "internal"
  kubernetes_cluster_id = "pike"
  vm_size               = "Standard_DS2_v2"
  node_count            = 1

  tags = {
    Environment = "Production"
  }
}
