resource "azurerm_tenant_template_deployment" "pike_gen" {
  name                     = "example"
  location                 = "West Europe"
  template_spec_version_id = "pike"
}
