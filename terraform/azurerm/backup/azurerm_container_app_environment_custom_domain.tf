resource "azurerm_container_app_environment_custom_domain" "pike_gen" {
  container_app_environment_id = "pike"
  certificate_blob_base64      = filebase64("testacc.pfx")
  dns_suffix                   = "acceptancetest.contoso.com"
}
