data "google_iap_tunnel_instance_iam_policy" "pike" {
  instance = "pike"
  zone     = "us-central1"
}

output "google_iap_tunnel_instance_iam_policy" {
  value = data.google_iap_tunnel_instance_iam_policy.pike
}
