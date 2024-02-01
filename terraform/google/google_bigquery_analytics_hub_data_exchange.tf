resource "google_bigquery_analytics_hub_data_exchange" "pike" {
  location         = "EUROPE"
  data_exchange_id = "my_data_exchange"
  display_name     = "my_data_exchange"
  description      = "example data exchange"
}
