resource "azurerm_cdn_frontdoor_secret" "pike_gen" {
  name                     = "example-customer-managed-secret"
  cdn_frontdoor_profile_id = "pike"

  secret {
    customer_certificate {
      key_vault_certificate_id = "pike"
    }
  }
}
