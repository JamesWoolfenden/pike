data "google_bigquery_table_iam_policy" "pike" {
  dataset_id = "pike"
  table_id   = "pike"
}
