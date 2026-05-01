resource "google_service_account_key" "pike" {
  service_account_id = "test123@pike-361314.iam.gserviceaccount.com"
}

output "google_service_account_key" {
  sensitive = true
  value     = google_service_account_key.pike
}
