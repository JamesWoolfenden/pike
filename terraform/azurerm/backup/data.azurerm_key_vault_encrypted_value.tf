data "azurerm_key_vault_encrypted_value" "pike" {
  algorithm        = "RSA1_5"
  key_vault_key_id = "https://pike-perm.vault.azure.net/keys/pike/af4f204b5e394fce8d98b26ae90c1d7b"
  encrypted_data   = "true"
}

output "value" {
  value = data.azurerm_key_vault_encrypted_value.pike
}
