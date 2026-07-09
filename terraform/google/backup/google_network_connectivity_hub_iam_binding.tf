resource "google_network_connectivity_hub_iam_binding" "pike" {
  hub = google_network_connectivity_hub.pike.name

  role = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
