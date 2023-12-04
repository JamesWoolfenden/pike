data "google_iap_tunnel_instance_iam_policy" "pike" {
  instance = "pike"
  zone     = "us-central1"
}

output "instance_policy" {
  value = data.google_iap_tunnel_instance_iam_policy.pike
}
