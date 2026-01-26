resource "google_discovery_engine_search_engine" "pike" {
  engine_id      = "example-engine-id"
  collection_id  = "default_collection"
  location       = google_discovery_engine_data_store.test_data_store.location
  display_name   = "Example Display Name"
  data_store_ids = [google_discovery_engine_data_store.test_data_store.data_store_id]
  search_engine_config {
  }
}
