resource "azurerm_linux_web_app_slot" "pike_gen" {
  name           = "example-slot"
  app_service_id = "pike"

  site_config {}
}
