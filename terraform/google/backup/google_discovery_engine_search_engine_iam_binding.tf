resource "google_discovery_engine_search_engine_iam_binding" "pike" {
  collection_id = google_discovery_engine_search_engine.pike.collection_id
  engine_id     = google_discovery_engine_search_engine.pike.engine_id
  location      = google_discovery_engine_search_engine.pike.location

  role = "roles/viewer"
  members = [
    "user:james.woolfenden@gmail.com",
  ]
}
