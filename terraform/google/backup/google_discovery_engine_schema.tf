resource "google_discovery_engine_schema" "pike" {
  location      = google_discovery_engine_data_store.test_data_store.location
  data_store_id = google_discovery_engine_data_store.test_data_store.data_store_id
  schema_id     = "schema-id"
  json_schema   = "{\"$schema\":\"https://json-schema.org/draft/2020-12/schema\",\"datetime_detection\":true,\"type\":\"object\",\"geolocation_detection\":true}"
}
