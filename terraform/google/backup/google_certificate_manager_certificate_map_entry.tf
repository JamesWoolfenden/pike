resource "google_certificate_manager_certificate_map" "certificate_map" {
  name        = "cert-map-entry"
  description = "My acceptance test certificate map"
  labels = {
    "terraform" : true,
    "acc-test" : true,
  }
}

resource "google_certificate_manager_certificate_map_entry" "default" {
  name        = "cert-map-entry"
  description = "My acceptance test certificate map entry"
  map         = google_certificate_manager_certificate_map.certificate_map.name
  labels = {
    "terraform" : true,
    "acc-test" : true,
  }
  certificates = [google_certificate_manager_certificate.certificate.id]
  matcher      = "PRIMARY"
}

resource "google_certificate_manager_certificate" "certificate" {
  name        = "cert-map-entry"
  description = "The default cert"
  scope       = "DEFAULT"
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


resource "google_certificate_manager_dns_authorization" "instance" {
  name        = "dns-auth"
  description = "The default dnss"
  domain      = "subdomain.hashicorptest.com"
}

resource "google_certificate_manager_dns_authorization" "instance2" {
  name        = "dns-auth2"
  description = "The default dnss"
  domain      = "subdomain2.hashicorptest.com"
}
