resource "azurerm_bot_connection" "pike_gen" {
  name                  = "example"
  bot_name              = "pike"
  location              = "pike"
  resource_group_name   = "pike"
  service_provider_name = "box"
  client_id             = "exampleId"
}
