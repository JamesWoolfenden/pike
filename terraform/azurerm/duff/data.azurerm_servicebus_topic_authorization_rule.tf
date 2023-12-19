data "azurerm_servicebus_topic_authorization_rule" "pike" {
  name           = "pike"
  namespace_name = "pike123456"
}
