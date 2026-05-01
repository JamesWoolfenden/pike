data "google_iap_tunnel_iam_policy" "pike" {}

output "google_iap_tunnel_iam_policy" {
  value = data.google_iap_tunnel_iam_policy.pike
}
