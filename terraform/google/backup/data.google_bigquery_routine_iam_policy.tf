data "google_bigquery_routine_iam_policy" "pike" {
  dataset_id = "pike"
  routine_id = "pike"
}

output "google_bigquery_routine_iam_policy" {
  value = data.google_bigquery_routine_iam_policy.pike
}
