data "google_certificate_manager_certificates" "pike" {
}

output "google_certificate_manager_certificates" {
  value = data.google_certificate_manager_certificates.pike
}
