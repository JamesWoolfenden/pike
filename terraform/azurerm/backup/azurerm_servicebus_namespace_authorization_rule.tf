resource "azurerm_servicebus_namespace_authorization_rule" "pike_gen" {
  name         = "examplerule"
  namespace_id = "pike"

  listen = true
  send   = true
  manage = false
}
