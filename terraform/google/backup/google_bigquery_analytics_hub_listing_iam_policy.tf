data "google_iam_policy" "admin4" {
  binding {
    role = "roles/viewer"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_bigquery_analytics_hub_listing_iam_policy" "policy" {
  data_exchange_id = google_bigquery_analytics_hub_listing.pike.data_exchange_id
  listing_id       = google_bigquery_analytics_hub_listing.pike.listing_id
  policy_data      = data.google_iam_policy.admin4.policy_data
}
