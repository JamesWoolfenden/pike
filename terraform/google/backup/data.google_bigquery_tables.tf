data "google_bigquery_tables" "pike" {
  dataset_id = "pike"
}

output "google_bigquery_tables" {
  value = data.google_bigquery_tables.pike
}
