data "azurerm_key_vault_access_policy" "pike" {
  name = "Key Management"
}

output "policy" {
  value = data.azurerm_key_vault_access_policy.pike
}
