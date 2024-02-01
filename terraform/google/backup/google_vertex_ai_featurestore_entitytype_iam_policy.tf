resource "google_vertex_ai_featurestore_entitytype_iam_policy" "pike" {
  provider     = google-beta
  featurestore = google_vertex_ai_featurestore.pike.name
  entitytype   = google_vertex_ai_featurestore_entitytype.pike.name
  policy_data  = data.google_iam_policy.admin2.policy_data
}

data "google_iam_policy" "admin2" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}
