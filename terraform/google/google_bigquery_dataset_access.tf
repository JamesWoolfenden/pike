resource "google_bigquery_dataset_access" "pike" {
  dataset_id    = google_bigquery_dataset.default.dataset_id
  role          = "OWNER"
  user_by_email = google_service_account.bqowner.email
}

resource "google_service_account" "bqowner" {
  account_id = "bqowner"
}
