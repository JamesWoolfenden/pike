data "google_bigquery_datapolicyv2_data_policy_iam_policy" "pike" {
  data_policy_id = "pike"
}

output "google_bigquery_datapolicyv2_data_policy_iam_policy" {
  value = data.google_bigquery_datapolicyv2_data_policy_iam_policy.pike
}
