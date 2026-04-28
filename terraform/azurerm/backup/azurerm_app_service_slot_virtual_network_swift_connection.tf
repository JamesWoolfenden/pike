resource "azurerm_app_service_slot_virtual_network_swift_connection" "pike_gen" {
  slot_name      = "pike"
  app_service_id = "pike"
  subnet_id      = "pike"
}
