resource "google_bigquery_connection" "pike" {
  connection_id = "my-connection"
  location      = "US"
  friendly_name = "something"
  description   = "a riveting description"
  cloud_resource {}
}
