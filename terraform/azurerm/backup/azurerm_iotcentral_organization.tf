resource "azurerm_iotcentral_organization" "pike_gen" {
  iotcentral_application_id = "pike"
  organization_id           = "example-parent-organization-id"
  display_name              = "Org example parent"
}
