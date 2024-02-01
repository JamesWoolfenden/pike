resource "google_bigquery_datapolicy_data_policy_iam_member" "pike" {
  data_policy_id = google_bigquery_datapolicy_data_policy.pike.data_policy_id
  role           = "roles/viewer"
  member         = "user:crwoolfenden@gmail.com"
}
