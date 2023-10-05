data "google_vertex_ai_featurestore_iam_policy" "pike" {
  provider     = google-beta
  featurestore = "projects/pike-gcp/locations/europe-west3/featurestores/pike"
}
