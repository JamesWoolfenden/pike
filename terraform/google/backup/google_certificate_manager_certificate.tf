resource "google_certificate_manager_certificate" "default" {
  name        = "dns-cert"
  description = "The default cert"
  scope       = "EDGE_CACHE"
  labels = {
    env = "test"
  }
  managed {
    domains = [
      google_certificate_manager_dns_authorization.instance.domain,
      google_certificate_manager_dns_authorization.instance2.domain,
    ]
    dns_authorizations = [
      google_certificate_manager_dns_authorization.instance.id,
      google_certificate_manager_dns_authorization.instance2.id,
    ]
  }
}
