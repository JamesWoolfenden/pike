data "google_service_account_id_token" "oidc" {
  target_audience = "https://foo.bar/"
}

output "oidc_token" {
  value     = data.google_service_account_id_token.oidc.id_token
  sensitive = true
}
