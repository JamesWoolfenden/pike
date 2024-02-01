data "google_iam_policy" "admin5" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_bigquery_analytics_hub_data_exchange_iam_policy" "policy" {
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.pike.data_exchange_id
  policy_data      = data.google_iam_policy.admin5.policy_data
}
