resource "azurerm_servicebus_subscription" "pike_gen" {
  name               = "tfex_servicebus_subscription"
  topic_id           = "pike"
  max_delivery_count = 1
}
