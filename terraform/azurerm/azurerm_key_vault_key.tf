resource "azurerm_key_vault_key" "pike" {
  name         = "generated-certificate"
  key_vault_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.KeyVault/vaults/pike2"
  key_type     = "RSA"
  key_size     = 2048

  key_opts = [
    "decrypt",
    "encrypt",
    "sign",
    "unwrapKey",
    "verify",
    "wrapKey",
  ]
}
