data "google_network_connectivity_hub_iam_policy" "pike" {
  hub = "pike"
}

output "google_network_connectivity_hub_iam_policy" {
  value = data.google_network_connectivity_hub_iam_policy.pike
}
