resource "google_certificate_manager_dns_authorization" "pike" {
  name        = "dns-auth"
  location    = "global"
  description = "The default dns"
  domain      = "subdomain.hashicorptest.com"
}
