resource "google_bigquery_datapolicy_data_policy_iam_binding" "pike" {
  data_policy_id = google_bigquery_datapolicy_data_policy.pike.data_policy_id
  role           = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
