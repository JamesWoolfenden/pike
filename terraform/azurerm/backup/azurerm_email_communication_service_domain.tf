resource "azurerm_email_communication_service_domain" "pike_gen" {
  name             = "AzureManagedDomain"
  email_service_id = "pike"

  domain_management = "AzureManaged"
}
