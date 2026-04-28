resource "azurerm_attestation_provider" "pike_gen" {
  name                = "exampleprovider"
  resource_group_name = "pike"
  location            = "pike"

  policy_signing_certificate_data = file("./example/cert.pem")
}
