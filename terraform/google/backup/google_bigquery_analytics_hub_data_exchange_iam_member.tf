resource "google_bigquery_analytics_hub_data_exchange_iam_member" "pike" {
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.pike.data_exchange_id
  role             = "roles/viewer"
  member           = "user:crwoolfenden@gmail.com"
}
