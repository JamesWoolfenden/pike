resource "azurerm_key_vault_managed_storage_account_sas_token_definition" "pike_gen" {
  name                       = "examplesasdefinition"
  validity_period            = "P1D"
  managed_storage_account_id = "pike"
  sas_template_uri           = "pike"
  sas_type                   = "account"
}
