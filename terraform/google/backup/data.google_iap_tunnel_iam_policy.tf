data "google_iap_tunnel_iam_policy" "pike" {}

output "tunnel" {
  value = data.google_iap_tunnel_iam_policy.pike
}
