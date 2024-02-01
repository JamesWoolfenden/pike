resource "google_vertex_ai_featurestore_iam_member" "pike" {
  provider     = google-beta
  featurestore = google_vertex_ai_featurestore.pike.name
  role         = "roles/viewer"
  member       = "user:crwoolfenden@gmail.com"
}
