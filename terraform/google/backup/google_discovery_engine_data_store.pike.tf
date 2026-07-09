resource "google_discovery_engine_data_store" "pike" {
  data_store_id     = "pike2"
  display_name      = "pike2"
  location          = "global"
  industry_vertical = "GENERIC"
  content_config    = "PUBLIC_WEBSITE"
  solution_types    = ["SOLUTION_TYPE_SEARCH"]
}
