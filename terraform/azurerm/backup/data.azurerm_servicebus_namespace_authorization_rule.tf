data "azurerm_servicebus_namespace_authorization_rule" "pike" {
  name         = "pike"
  namespace_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.ServiceBus/namespaces/pike"
}
