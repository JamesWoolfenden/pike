data "google_service_account" "default" {
  account_id = "pike-service@examplea.iam.gserviceaccount.com"
}

output "service_account" {
  value = data.google_service_account.default
}
