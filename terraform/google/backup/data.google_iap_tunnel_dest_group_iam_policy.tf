data "google_iap_tunnel_dest_group_iam_policy" "pike" {
  dest_group = "pike"
}

output "google_iap_tunnel_dest_group_iam_policy" {
  value = data.google_iap_tunnel_dest_group_iam_policy.pike
}
