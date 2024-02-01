resource "google_vertex_ai_endpoint_iam_binding" "pike" {
  provider = google-beta
  endpoint = google_vertex_ai_endpoint.pike.name
  role     = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
