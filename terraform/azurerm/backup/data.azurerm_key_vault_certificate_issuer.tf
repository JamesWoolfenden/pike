data "azurerm_key_vault_certificate_issuer" "pike" {
  name         = "jimbo"
  key_vault_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.KeyVault/vaults/pike-perm"
}
