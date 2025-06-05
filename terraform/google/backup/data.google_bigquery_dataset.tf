data "google_bigquery_dataset" "pike" {
  dataset_id = "my-bq-dataset"
  project    = "pike"
}
