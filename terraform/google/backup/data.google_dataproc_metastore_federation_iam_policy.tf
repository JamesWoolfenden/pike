data "google_dataproc_metastore_federation_iam_policy" "pike" {
  provider      = google-beta
  federation_id = "pike"
}
