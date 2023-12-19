data "azurerm_servicebus_queue_authorization_rule" "pike" {
  name                = "pike"
  resource_group_name = "pike"
}
