resource "azurerm_logic_app_integration_account_agreement" "pike_gen" {
  name                     = "example-agreement"
  resource_group_name      = "pike"
  integration_account_name = "pike"
  agreement_type           = "AS2"
  host_partner_name        = "pike"
  guest_partner_name       = "pike"
  content                  = file("testdata/integration_account_agreement_content_as2.json")

  host_identity {
    qualifier = "AS2Identity"
    value     = "FabrikamNY"
  }

  guest_identity {
    qualifier = "AS2Identity"
    value     = "FabrikamDC"
  }
}
