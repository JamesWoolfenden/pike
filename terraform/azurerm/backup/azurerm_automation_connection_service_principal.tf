resource "azurerm_automation_connection_service_principal" "pike_gen" {
  name                    = "connection-example"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  application_id          = "00000000-0000-0000-0000-000000000000"
  tenant_id               = "pike"
  subscription_id         = "pike"
  certificate_thumbprint  = file("automation_certificate_test.thumb")
}
