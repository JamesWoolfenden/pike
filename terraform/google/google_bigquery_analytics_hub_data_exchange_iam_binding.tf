resource "google_bigquery_analytics_hub_data_exchange_iam_binding" "pike" {
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.pike.data_exchange_id
  role             = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
