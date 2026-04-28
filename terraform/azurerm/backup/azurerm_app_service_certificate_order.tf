resource "azurerm_app_service_certificate_order" "pike_gen" {
  name                = "example-cert-order"
  resource_group_name = "pike"
  location            = "global"
  distinguished_name  = "CN=example.com"
  product_type        = "Standard"
}
