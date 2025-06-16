data "google_bigquery_table" "pike" {
  table_id   = "pike"
  dataset_id = "pike"
}

output "google_bigquery_table" {
  value = data.google_bigquery_table.pike
}
