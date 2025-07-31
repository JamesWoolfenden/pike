data "google_site_verification_token" "pike" {
  type                = "SITE"
  verification_method = "META"
  identifier          = "https://www.example.com"
}

output "google_site_verification_token" {
  value = data.google_site_verification_token.pike
}
