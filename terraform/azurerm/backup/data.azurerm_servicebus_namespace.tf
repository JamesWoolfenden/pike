data "azurerm_servicebus_namespace" "pike" {
  resource_group_name = "pike"
  name                = "pike"
}
