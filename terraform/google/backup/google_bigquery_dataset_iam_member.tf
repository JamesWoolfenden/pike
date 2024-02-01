resource "google_bigquery_dataset_iam_member" "pike" {
  dataset_id = google_bigquery_dataset.default.dataset_id
  role       = "roles/bigquery.dataEditor"

  member = "user:crwoolfenden@gmail.com"
}
