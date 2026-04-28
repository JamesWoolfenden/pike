resource "azurerm_api_management_notification_recipient_user" "pike_gen" {
  api_management_id = "pike"
  notification_type = "AccountClosedPublisher"
  user_id           = "pike"
}
