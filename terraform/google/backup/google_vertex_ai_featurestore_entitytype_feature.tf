resource "google_vertex_ai_featurestore_entitytype_feature" "pike" {
  provider   = google-beta
  entitytype = google_vertex_ai_featurestore_entitytype.pike.name

  name = "terraform"
  labels = {
    foo = "bar"
  }

  value_type = "INT64_ARRAY"
}
