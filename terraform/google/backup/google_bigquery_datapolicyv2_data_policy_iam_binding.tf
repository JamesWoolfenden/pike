resource "google_bigquery_datapolicyv2_data_policy_iam_binding" "pike" {
  project        = google_bigquery_datapolicyv2_data_policy.pike.project
  location       = google_bigquery_datapolicyv2_data_policy.pike.location
  data_policy_id = google_bigquery_datapolicyv2_data_policy.pike.data_policy_id
  role           = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
