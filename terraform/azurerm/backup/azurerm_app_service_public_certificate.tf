resource "azurerm_app_service_public_certificate" "pike_gen" {
  resource_group_name  = "pike"
  app_service_name     = "pike"
  certificate_name     = "example-public-certificate"
  certificate_location = "Unknown"
  blob                 = filebase64("app_service_public_certificate.cer")
}
