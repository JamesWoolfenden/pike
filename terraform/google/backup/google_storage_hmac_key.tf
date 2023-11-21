resource "google_storage_hmac_key" "pike" {
  service_account_email = "pike-service@pike-gcp.iam.gserviceaccount.com"
}
