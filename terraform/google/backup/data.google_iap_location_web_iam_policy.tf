data "google_iap_location_web_iam_policy" "pike" {
  location = "us-central1"
}

output "google_iap_location_web_iam_policy" {
  value = data.google_iap_location_web_iam_policy.pike
}
