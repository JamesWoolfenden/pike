resource "google_discovery_engine_recommendation_engine" "pike" {
  engine_id         = "recommendation-engine-id"
  location          = google_discovery_engine_data_store.test_data_store.location
  display_name      = "Example Recommendation Engine"
  data_store_ids    = [google_discovery_engine_data_store.test_data_store.data_store_id]
  industry_vertical = "GENERIC"
  common_config {
    company_name = "test-company"
  }
}
