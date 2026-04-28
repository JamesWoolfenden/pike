resource "azurerm_cdn_frontdoor_custom_domain" "pike_gen" {
  name                     = "example-customDomain"
  cdn_frontdoor_profile_id = "pike"
  dns_zone_id              = "pike"
  host_name                = "contoso.fabrikam.com"

  tls {
    certificate_type    = "ManagedCertificate"
    minimum_tls_version = "TLS12"
  }
}
