resource "azurerm_relay_hybrid_connection_authorization_rule" "pike_gen" {
  name                   = "example"
  resource_group_name    = "pike"
  hybrid_connection_name = "pike"
  namespace_name         = "pike"


  listen = true
  send   = true
  manage = false
}
