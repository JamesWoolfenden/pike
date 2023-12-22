data "google_vertex_ai_endpoint_iam_policy" "pike" {
  provider = google-beta
  project  = "pike-gcp"
  location = "europe-west2"
  endpoint = "pike"
}
