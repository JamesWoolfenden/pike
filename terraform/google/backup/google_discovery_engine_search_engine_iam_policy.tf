data "google_iam_policy" "google_discovery_engine_search_engine" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_discovery_engine_search_engine_iam_policy" "pike" {
  collection_id = google_discovery_engine_search_engine.pike.collection_id
  engine_id     = google_discovery_engine_search_engine.pike.engine_id
  location      = google_discovery_engine_search_engine.pike.location

  policy_data = data.google_iam_policy.google_discovery_engine_search_engine.policy_data
}
