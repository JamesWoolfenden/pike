resource "azurerm_app_service_virtual_network_swift_connection" "pike_gen" {
  app_service_id = "pike"
  subnet_id      = "pike"
}
