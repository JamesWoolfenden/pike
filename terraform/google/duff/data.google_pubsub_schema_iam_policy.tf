data "google_pubsub_schema_iam_policy" "pike" {
  provider = google-beta
  schema   = "pike"
  project  = "pike-gcp"
}
