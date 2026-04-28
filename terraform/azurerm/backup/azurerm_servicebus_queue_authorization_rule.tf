resource "azurerm_servicebus_queue_authorization_rule" "pike_gen" {
  name     = "examplerule"
  queue_id = "pike"

  listen = true
  send   = true
  manage = false
}
