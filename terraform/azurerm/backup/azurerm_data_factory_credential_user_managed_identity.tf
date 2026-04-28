resource "azurerm_data_factory_credential_user_managed_identity" "pike_gen" {
  name            = "pike"
  description     = "Short description of this credential"
  data_factory_id = "pike"
  identity_id     = "pike"

  annotations = ["example", "example2"]
}
