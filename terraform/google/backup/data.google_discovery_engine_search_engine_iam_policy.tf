data "google_discovery_engine_search_engine_iam_policy" "pike" {
  collection_id = "default_collection"
  engine_id     = "example-engine-id"
  location      = "global"
}

output "google_discovery_engine_search_engine_iam_policy" {
  value = data.google_discovery_engine_search_engine_iam_policy.pike
}
