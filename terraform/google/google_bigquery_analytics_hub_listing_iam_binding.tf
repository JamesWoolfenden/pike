resource "google_bigquery_analytics_hub_listing_iam_binding" "pike" {

  data_exchange_id = google_bigquery_analytics_hub_listing.pike.data_exchange_id
  listing_id       = google_bigquery_analytics_hub_listing.pike.listing_id
  role             = "roles/viewer"
  members = [
    "user:crwoolfenden@gmail.com",
  ]
}
