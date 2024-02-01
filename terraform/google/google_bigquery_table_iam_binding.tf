resource "google_bigquery_table_iam_binding" "pike" {
  dataset_id = google_bigquery_dataset.default.id
  table_id   = google_bigquery_table.default.id
  role       = "roles/bigquery.dataOwner"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
