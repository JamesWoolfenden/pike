resource "google_vertex_ai_feature_online_store" "pike" {
  name = "example_feature_online_store"
  labels = {
    foo = "bar"
  }
  region = "us-central1"
  bigtable {
    auto_scaling {
      min_node_count         = 1
      max_node_count         = 3
      cpu_utilization_target = 50
    }
  }
}
