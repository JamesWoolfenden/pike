data "azurerm_kubernetes_cluster" "pike_gen" {
  name                = "myakscluster"
  resource_group_name = "my-example-resource-group"
}
