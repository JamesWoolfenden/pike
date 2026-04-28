resource "azurerm_gallery_application" "pike_gen" {
  name              = "example-app"
  gallery_id        = "pike"
  location          = "pike"
  supported_os_type = "Linux"
}
