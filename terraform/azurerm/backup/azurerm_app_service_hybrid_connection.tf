resource "azurerm_app_service_hybrid_connection" "pike_gen" {
  app_service_name    = "pike"
  resource_group_name = "pike"
  relay_id            = "pike"
  hostname            = "testhostname.example"
  port                = 8080
  send_key_name       = "exampleSharedAccessKey"
}
