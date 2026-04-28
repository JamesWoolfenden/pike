resource "azurerm_notification_hub_authorization_rule" "pike_gen" {
  name                  = "management-auth-rule"
  notification_hub_name = "pike"
  namespace_name        = "pike"
  resource_group_name   = "pike"
  manage                = true
  send                  = true
  listen                = true
}
