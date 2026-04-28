data "azurerm_kubernetes_cluster_node_pool" "pike_gen" {
  name                    = "existing"
  kubernetes_cluster_name = "existing-cluster"
  resource_group_name     = "existing-resource-group"
}
