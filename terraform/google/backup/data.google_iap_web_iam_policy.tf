data "google_iap_web_iam_policy" "pike" {}

output "policy4" {
  value = data.google_iap_web_iam_policy.pike
}
