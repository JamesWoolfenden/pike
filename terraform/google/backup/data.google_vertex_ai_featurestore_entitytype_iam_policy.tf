data "google_vertex_ai_featurestore_entitytype_iam_policy" "pike" {
  provider     = google-beta
  featurestore = "projects/pike-gcp/locations/europe-west3/featurestores/pike"
  entitytype   = "projects/pike-gcp/locations/us-central1/featurestores/pike/entityTypes/pike"
}
