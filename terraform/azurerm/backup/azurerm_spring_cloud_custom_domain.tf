resource "azurerm_spring_cloud_custom_domain" "pike_gen" {
  name                = join(".", ["pike", "pike"])
  spring_cloud_app_id = "pike"
}
