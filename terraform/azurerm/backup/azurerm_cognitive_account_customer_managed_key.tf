resource "azurerm_cognitive_account_customer_managed_key" "pike_gen" {
  cognitive_account_id = "pike"
  key_vault_key_id     = "pike"
  identity_client_id   = "pike"
}
