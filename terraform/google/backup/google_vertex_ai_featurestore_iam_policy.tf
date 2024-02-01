resource "google_vertex_ai_featurestore_iam_policy" "pike" {
  provider     = google-beta
  featurestore = google_vertex_ai_featurestore.pike.name
  policy_data  = data.google_iam_policy.admin.policy_data
}


data "google_iam_policy" "admin" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}
