resource "azurerm_relay_namespace_authorization_rule" "pike_gen" {
  name                = "example"
  resource_group_name = "pike"
  namespace_name      = "pike"

  listen = true
  send   = true
  manage = false
}
