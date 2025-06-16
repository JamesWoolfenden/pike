data "google_dataproc_metastore_database_iam_policy" "pike" {
  database   = "pike"
  service_id = "pike"
}

output "google_dataproc_metastore_database_iam_policy" {
  value = data.google_dataproc_metastore_database_iam_policy.pike
}
