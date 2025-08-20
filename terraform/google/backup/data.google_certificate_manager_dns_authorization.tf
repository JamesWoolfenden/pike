data "google_certificate_manager_dns_authorization" "pike" {
}

output "google_certificate_manager_dns_authorization" {
  value = data.google_certificate_manager_dns_authorization.pike
}
