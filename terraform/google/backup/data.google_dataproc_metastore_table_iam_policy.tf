data "google_dataproc_metastore_table_iam_policy" "pike" {
  database_id = "pike"
  table       = "pike"
  service_id  = "pike"

}

output "google_dataproc_metastore_table_iam_policy" {
  value = data.google_dataproc_metastore_table_iam_policy.pike
}
