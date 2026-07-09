data "google_iam_policy" "google_network_connectivity_hub" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_network_connectivity_hub_iam_policy" "pike" {
  hub = google_network_connectivity_hub.pike.name

  policy_data = data.google_iam_policy.google_network_connectivity_hub.policy_data
}
