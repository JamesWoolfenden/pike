data "google_iam_policy" "admin4" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.Woolfenden@gmail.com",
    ]
  }
}

resource "google_bigquery_datapolicy_data_policy_iam_policy" "policy" {
  project        = google_bigquery_datapolicy_data_policy.pike.project
  location       = google_bigquery_datapolicy_data_policy.pike.location
  data_policy_id = google_bigquery_datapolicy_data_policy.pike.data_policy_id
  policy_data    = data.google_iam_policy.admin4.policy_data
}
