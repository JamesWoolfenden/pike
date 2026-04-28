resource "azurerm_spring_cloud_app" "pike_gen" {
  name                = "example-springcloudapp"
  resource_group_name = "pike"
  service_name        = "pike"

  identity {
    type = "SystemAssigned"
  }
}
