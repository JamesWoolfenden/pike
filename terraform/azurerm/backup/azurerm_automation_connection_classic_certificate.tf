resource "azurerm_automation_connection_classic_certificate" "pike_gen" {
  name                    = "connection-example"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  certificate_asset_name  = "cert1"
  subscription_name       = "subs1"
  subscription_id         = "pike"
}
