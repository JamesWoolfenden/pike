resource "azurerm_automation_runtime_environment_package" "pike_gen" {
  name                              = "example-package"
  automation_runtime_environment_id = "pike"
  content_uri                       = "https://www.powershellgallery.com/api/v2/package/example-package/1.0.0"
}
