data "google_iam_policy" "admin3" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_bigquery_datapolicy_data_policy_iam_policy" "policy" {
  data_policy_id = google_bigquery_datapolicy_data_policy.pike.data_policy_id
  policy_data    = data.google_iam_policy.admin3.policy_data
}
