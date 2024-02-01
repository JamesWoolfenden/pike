resource "google_bigquery_dataset_iam_binding" "pike" {
  dataset_id = google_bigquery_dataset.default.dataset_id
  role       = "roles/bigquery.dataViewer"

  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
