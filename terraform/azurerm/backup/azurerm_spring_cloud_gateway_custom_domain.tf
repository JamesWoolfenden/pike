resource "azurerm_spring_cloud_gateway_custom_domain" "pike_gen" {
  name                    = "example.com"
  spring_cloud_gateway_id = "pike"
}
