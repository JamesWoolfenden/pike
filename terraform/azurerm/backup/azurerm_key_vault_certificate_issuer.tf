resource "azurerm_key_vault_certificate_issuer" "pike_gen" {
  name          = "example-issuer"
  org_id        = "ExampleOrgName"
  key_vault_id  = "pike"
  provider_name = "DigiCert"
  account_id    = "0000"
}
