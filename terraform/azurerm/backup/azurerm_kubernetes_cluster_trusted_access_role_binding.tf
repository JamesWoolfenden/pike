resource "azurerm_kubernetes_cluster_trusted_access_role_binding" "pike_gen" {
  kubernetes_cluster_id = "pike"
  name                  = "example"
  roles                 = "example-value"
  source_resource_id    = "pike"
}
