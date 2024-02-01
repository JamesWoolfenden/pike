resource "google_vertex_ai_featurestore_entitytype_iam_member" "pike" {
  provider     = google-beta
  featurestore = google_vertex_ai_featurestore.pike.name
  entitytype   = google_vertex_ai_featurestore_entitytype.pike.name
  role         = "roles/viewer"
  member       = "user:crwoolfenden@gmail.com"
}
