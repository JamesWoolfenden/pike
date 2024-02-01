resource "google_bigquery_table_iam_member" "pike" {
  dataset_id = google_bigquery_dataset.default.id
  table_id   = google_bigquery_table.default.id
  role       = "roles/bigquery.dataOwner"
  member     = "user:crwoolfenden@gmail.com"
}
