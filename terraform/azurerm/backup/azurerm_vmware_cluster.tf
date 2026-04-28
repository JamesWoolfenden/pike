resource "azurerm_vmware_cluster" "pike_gen" {
  name               = "example-Cluster"
  vmware_cloud_id    = "pike"
  cluster_node_count = 3
  sku_name           = "av36"
}
