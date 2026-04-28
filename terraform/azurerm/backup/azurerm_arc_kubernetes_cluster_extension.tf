resource "azurerm_arc_kubernetes_cluster_extension" "pike_gen" {
  name           = "example-ext"
  cluster_id     = "pike"
  extension_type = "microsoft.flux"
}
