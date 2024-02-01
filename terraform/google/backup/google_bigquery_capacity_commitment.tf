resource "google_bigquery_capacity_commitment" "pike" {
  capacity_commitment_id = "example-commitment"

  location   = "europe-west2"
  slot_count = 100
  plan       = "FLEX_FLAT_RATE"
  edition    = "ENTERPRISE"
}
