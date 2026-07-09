resource "google_discovery_engine_search_engine" "pike" {
  collection_id  = "default_collection"
  engine_id      = "example-engine-id"
  display_name   = "pike"
  location       = "global"
  data_store_ids = [google_discovery_engine_data_store.pike.data_store_id]

  search_engine_config {}
}
