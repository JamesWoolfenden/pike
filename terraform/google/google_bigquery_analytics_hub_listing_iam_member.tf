resource "google_bigquery_analytics_hub_listing_iam_member" "pike" {
  data_exchange_id = google_bigquery_analytics_hub_listing.pike.data_exchange_id
  listing_id       = google_bigquery_analytics_hub_listing.pike.listing_id
  role             = "roles/viewer"
  member           = "user:crwoolfenden@gmail.com"
}
