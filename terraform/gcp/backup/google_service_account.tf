resource "google_service_account" "pike" {
  account_id   = "pike-service-account-id"
  display_name = "Service Account update"
}

output "service_account_pike" {
  value = google_service_account.pike
}
