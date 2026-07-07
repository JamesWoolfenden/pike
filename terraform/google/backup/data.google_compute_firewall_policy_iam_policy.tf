data "google_compute_firewall_policy_iam_policy" "pike" {
  name = "pike"
}

output "google_compute_firewall_policy_iam_policy" {
  value = data.google_compute_firewall_policy_iam_policy.pike
}
