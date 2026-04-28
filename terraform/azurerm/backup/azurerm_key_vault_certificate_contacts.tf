resource "azurerm_key_vault_certificate_contacts" "pike_gen" {
  key_vault_id = "pike"

  contact {
    email = "example@example.com"
    name  = "example"
    phone = "01234567890"
  }

  contact {
    email = "example2@example.com"
  }

  depends_on = [
    azurerm_key_vault_access_policy.example
  ]
}
