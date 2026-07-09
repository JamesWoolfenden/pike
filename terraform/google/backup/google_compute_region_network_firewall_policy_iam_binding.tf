resource "google_compute_region_network_firewall_policy_iam_binding" "pike" {
  region = google_compute_region_network_firewall_policy.pike.region
  name   = google_compute_region_network_firewall_policy.pike.name

  role = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
