resource "google_vertex_ai_endpoint_iam_policy" "pike" {
  provider    = google-beta
  endpoint    = google_vertex_ai_endpoint.pike.name
  policy_data = data.google_iam_policy.admin3.policy_data
}

data "google_iam_policy" "admin3" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}
