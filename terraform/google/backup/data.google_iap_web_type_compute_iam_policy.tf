data "google_iap_web_type_compute_iam_policy" "pike" {}

output "policy" {
  value = data.google_iap_web_type_compute_iam_policy.pike
}
