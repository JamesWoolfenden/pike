resource "azurerm_notification_hub" "pike_gen" {
  name                = "mynotificationhub"
  namespace_name      = "pike"
  resource_group_name = "pike"
  location            = "pike"
}
