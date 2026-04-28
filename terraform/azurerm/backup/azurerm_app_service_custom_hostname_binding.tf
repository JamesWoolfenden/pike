resource "azurerm_app_service_custom_hostname_binding" "pike_gen" {
  hostname            = "www.mywebsite.com"
  app_service_name    = "pike"
  resource_group_name = "pike"
}
