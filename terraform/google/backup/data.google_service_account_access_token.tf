data "google_service_account_access_token" "default" {
  provider               = google
  target_service_account = "pike-service@pike-gcp.iam.gserviceaccount.com"
  scopes                 = ["userinfo-email", "cloud-platform"]
  lifetime               = "300s"
}
