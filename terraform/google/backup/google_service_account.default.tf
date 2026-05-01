data "google_service_account" "default" {
  account_id = "pike-service@examplea.iam.gserviceaccount.com"
}

output "google_service_account.default" {
  value = data.google_service_account.default
}
