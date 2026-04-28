resource "azurerm_eventhub_authorization_rule" "pike_gen" {
  name                = "navi"
  namespace_name      = "pike"
  eventhub_name       = "pike"
  resource_group_name = "pike"
  listen              = true
  send                = false
  manage              = false
}
