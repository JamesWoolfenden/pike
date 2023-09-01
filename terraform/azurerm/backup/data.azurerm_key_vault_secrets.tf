data "azurerm_key_vault_secrets" "pike" {
  key_vault_id = "/subscriptions/037ce662-dfc1-4b8b-a8a7-6c414b540ed6/resourceGroups/pike/providers/Microsoft.KeyVault/vaults/pike-perm"
}

output "secrets" {
  value = data.azurerm_key_vault_secrets.pike
}
