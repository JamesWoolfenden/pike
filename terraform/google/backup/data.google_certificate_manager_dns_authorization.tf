data "google_certificate_manager_dns_authorization" "pike" {
  provider = google-beta
}

output "google_certificate_manager_dns_authorization" {
  value = data.google_certificate_manager_dns_authorization.pike
}
