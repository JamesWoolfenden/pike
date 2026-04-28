data "azurerm_notification_hub_namespace" "pike_gen" {
  name                = "my-namespace"
  resource_group_name = "my-resource-group"
}
