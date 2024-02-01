resource "google_bigquery_analytics_hub_listing" "pike" {
  location         = "EUROPE"
  data_exchange_id = google_bigquery_analytics_hub_data_exchange.pike.data_exchange_id
  listing_id       = "my_listing"
  display_name     = "my_listing"
  description      = "example data exchange"

  bigquery_dataset {
    dataset = google_bigquery_dataset.default.id
  }
}
