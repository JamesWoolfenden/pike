data "google_service_account_jwt" "pike" {
  target_service_account = "pike-service@pike-gcp.iam.gserviceaccount.com"

  payload = jsonencode({
    foo : "bar",
    sub : "subject",
  })

  expires_in = 60
}

output "jwt" {
  value     = data.google_service_account_jwt.pike.jwt
  sensitive = true
}
