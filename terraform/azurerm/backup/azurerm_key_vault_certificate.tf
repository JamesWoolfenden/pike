resource "azurerm_key_vault_certificate" "pike_gen" {
  name         = "imported-cert"
  key_vault_id = "pike"

  certificate {
    contents = filebase64("certificate-to-import.pfx")
  }
}
