resource "google_bigquery_dataset" "pike" {
  dataset_id                  = "foo"
  friendly_name               = "pike"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    pike = "permissions"
  }
}
