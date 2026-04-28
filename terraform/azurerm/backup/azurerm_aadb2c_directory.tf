resource "azurerm_aadb2c_directory" "pike_gen" {
  country_code            = "US"
  data_residency_location = "United States"
  display_name            = "example-b2c-tenant"
  domain_name             = "exampleb2ctenant.onmicrosoft.com"
  resource_group_name     = "example-rg"
  sku_name                = "PremiumP1"
}
