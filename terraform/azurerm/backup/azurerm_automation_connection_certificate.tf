resource "azurerm_automation_connection_certificate" "pike_gen" {
  name                        = "connection-example"
  resource_group_name         = "pike"
  automation_account_name     = "pike"
  automation_certificate_name = "pike"
  subscription_id             = "pike"
}
