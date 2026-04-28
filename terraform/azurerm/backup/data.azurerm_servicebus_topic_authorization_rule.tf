data "azurerm_servicebus_topic_authorization_rule" "pike_gen" {
  name                = "example-tfex_name"
  resource_group_name = "example-resources"
  namespace_name      = "example-namespace"
  topic_name          = "example-servicebus_topic"
}
