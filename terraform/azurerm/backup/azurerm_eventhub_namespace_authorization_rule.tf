resource "azurerm_eventhub_namespace_authorization_rule" "pike_gen" {
  name                = "navi"
  namespace_name      = "pike"
  resource_group_name = "pike"

  listen = true
  send   = false
  manage = false
}
