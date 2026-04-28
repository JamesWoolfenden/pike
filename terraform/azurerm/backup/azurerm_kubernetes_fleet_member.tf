resource "azurerm_kubernetes_fleet_member" "pike_gen" {
  kubernetes_cluster_id = "pike"
  kubernetes_fleet_id   = "pike"
  name                  = "example"
}
