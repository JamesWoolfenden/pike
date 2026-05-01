data "google_iap_web_iam_policy" "pike" {}

output "google_iap_web_iam_policy" {
  value = data.google_iap_web_iam_policy.pike
}
