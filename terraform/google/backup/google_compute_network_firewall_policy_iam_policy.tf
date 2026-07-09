data "google_iam_policy" "google_compute_network_firewall_policy" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_compute_network_firewall_policy_iam_policy" "pike" {
  name = google_compute_network_firewall_policy.pike.name

  policy_data = data.google_iam_policy.google_compute_network_firewall_policy.policy_data
}
