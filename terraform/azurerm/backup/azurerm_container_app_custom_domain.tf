resource "azurerm_container_app_custom_domain" "pike_gen" {
  name                                     = trimsuffix(trimprefix("pike", "asuid."), ".")
  container_app_id                         = "pike"
  container_app_environment_certificate_id = "pike"
  certificate_binding_type                 = "SniEnabled"
}
