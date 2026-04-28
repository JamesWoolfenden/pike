data "azurerm_container_app_environment_certificate" "pike_gen" {
  name                         = "mycertificate"
  container_app_environment_id = "pike"
}
