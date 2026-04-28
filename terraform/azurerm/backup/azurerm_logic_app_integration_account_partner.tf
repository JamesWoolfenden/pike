resource "azurerm_logic_app_integration_account_partner" "pike_gen" {
  name                     = "example-iap"
  resource_group_name      = "pike"
  integration_account_name = "pike"

  business_identity {
    qualifier = "ZZ"
    value     = "AA"
  }
}
