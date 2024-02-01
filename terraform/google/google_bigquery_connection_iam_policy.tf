resource "google_bigquery_connection_iam_policy" "pike" {
  connection_id = google_bigquery_connection.pike.connection_id
  policy_data   = data.google_iam_policy.admin.policy_data
}
