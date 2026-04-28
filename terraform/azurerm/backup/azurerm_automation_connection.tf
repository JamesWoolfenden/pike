resource "azurerm_automation_connection" "pike_gen" {
  name                    = "connection-example"
  resource_group_name     = "pike"
  automation_account_name = "pike"
  type                    = "AzureServicePrincipal"

  values = {
    "ApplicationId" : "00000000-0000-0000-0000-000000000000",
    "TenantId" : data.azurerm_client_config.example.tenant_id,
    "SubscriptionId" : data.azurerm_client_config.example.subscription_id,
    "CertificateThumbprint" : "sample-certificate-thumbprint",
  }
}
