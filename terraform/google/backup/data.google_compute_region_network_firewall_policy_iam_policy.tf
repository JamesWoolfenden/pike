data "google_compute_region_network_firewall_policy_iam_policy" "pike" {
  region = "us-central1"
  name   = "pike"
}

output "google_compute_region_network_firewall_policy_iam_policy" {
  value = data.google_compute_region_network_firewall_policy_iam_policy.pike
}
