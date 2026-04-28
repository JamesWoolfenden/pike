resource "azurerm_notification_hub_namespace" "pike_gen" {
  name                = "myappnamespace"
  resource_group_name = "pike"
  location            = "pike"
  namespace_type      = "NotificationHub"
  sku_name            = "Free"
}
