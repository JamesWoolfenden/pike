resource "azurerm_api_management_notification_recipient_email" "pike_gen" {
  api_management_id = "pike"
  notification_type = "AccountClosedPublisher"
  email             = "foo@bar.com"
}
