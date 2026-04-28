resource "azurerm_app_service_slot_custom_hostname_binding" "pike_gen" {
  app_service_slot_id = "pike"
  hostname            = "www.mywebsite.com"
}
