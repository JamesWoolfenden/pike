data "google_service_account_iam_policy" "pike" {
  service_account_id = google_service_account.myaccount.name
}

output "policy" {
  value = data.google_service_account_iam_policy.pike
}
