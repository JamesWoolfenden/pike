data "google_bigquery_datasets" "pike" {
}

output "google_bigquery_datasets" {
  value = data.google_bigquery_datasets.pike
}
