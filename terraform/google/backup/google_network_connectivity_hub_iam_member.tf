resource "google_network_connectivity_hub_iam_member" "pike" {
  hub = google_network_connectivity_hub.pike.name

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}
