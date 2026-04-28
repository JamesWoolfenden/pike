resource "azurerm_servicebus_topic_authorization_rule" "pike_gen" {
  name     = "tfex_servicebus_topic_sasPolicy"
  topic_id = "pike"
  listen   = true
  send     = false
  manage   = false
}
