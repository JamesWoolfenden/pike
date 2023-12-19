data "azurerm_spring_cloud_app" "pike" {
  service_name        = "pike"
  resource_group_name = "pike"
  name                = "pike"
}
