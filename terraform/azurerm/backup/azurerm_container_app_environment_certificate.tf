resource "azurerm_container_app_environment_certificate" "pike_gen" {
  name                         = "myfriendlyname"
  container_app_environment_id = "pike"
  certificate_blob_base64      = filebase64("path/to/certificate_file.pfx")
}
