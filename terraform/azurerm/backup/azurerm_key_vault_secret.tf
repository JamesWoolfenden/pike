resource "azurerm_key_vault_secret" "pike_gen" {
  name         = "secret-sauce"
  value        = "szechuan"
  key_vault_id = "pike"
}
