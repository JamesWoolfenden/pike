resource "google_compute_region_network_firewall_policy_iam_member" "pike" {
  region = google_compute_region_network_firewall_policy.pike.region
  name   = google_compute_region_network_firewall_policy.pike.name

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
