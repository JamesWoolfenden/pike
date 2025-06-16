data "google_privateca_certificate_authority" "pike" {
  location                 = "us-central1"
  pool                     = "pool-name"
  certificate_authority_id = "ca-id"
}

output "google_privateca_certificate_authority" {
  value = data.google_privateca_certificate_authority.pike
}
