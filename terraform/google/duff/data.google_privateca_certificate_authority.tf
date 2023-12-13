data "google_privateca_certificate_authority" "pike" {
  location = "us-central1"
}

output "auth" {
  value = data.google_privateca_certificate_authority.pike
}
