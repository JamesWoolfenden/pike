data "azurerm_notification_hub" "pike_gen" {
  name                = "notification-hub"
  namespace_name      = "namespace-name"
  resource_group_name = "resource-group-name"
}
