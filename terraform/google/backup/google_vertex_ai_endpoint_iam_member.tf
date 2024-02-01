resource "google_vertex_ai_endpoint_iam_member" "pike" {
  provider = google-beta
  endpoint = google_vertex_ai_endpoint.pike.name

  role   = "roles/viewer"
  member = "user:crwoolfenden@gmail.com"
}
