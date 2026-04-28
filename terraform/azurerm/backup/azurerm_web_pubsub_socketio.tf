resource "azurerm_web_pubsub_socketio" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  location            = "pike"
  sku                 = "Free_F1"
}
