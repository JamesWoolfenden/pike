resource "azurerm_extended_location_custom_location" "pike_gen" {
  name                = "example-custom-location"
  resource_group_name = "pike"
  location            = "West Europe"
  cluster_extension_ids = [
    azurerm_arc_kubernetes_cluster_extension.example.id
  ]
  display_name     = "example-custom-location"
  namespace        = "example-namespace"
  host_resource_id = "pike"
  authentication {
    value = base64encode("pike")
  }
}
