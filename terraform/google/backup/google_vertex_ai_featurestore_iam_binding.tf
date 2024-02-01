resource "google_vertex_ai_featurestore_iam_binding" "pike" {
  provider     = google-beta
  featurestore = google_vertex_ai_featurestore.pike.name
  role         = "roles/viewer"
  members      = ["user:crwoolfenden@gmail.com"]
}
