resource "google_vertex_ai_feature_group_feature" "pike" {
  name          = "example_feature"
  region        = "us-central1"
  feature_group = google_vertex_ai_feature_group.pike.name
  description   = "A sample feature"
  labels = {
    label-one = "value-one"
  }
}
