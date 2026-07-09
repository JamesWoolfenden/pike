resource "google_compute_network_firewall_policy_iam_member" "pike" {
  name = google_compute_network_firewall_policy.pike.name

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
