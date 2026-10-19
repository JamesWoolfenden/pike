resource "google_vertex_ai_tensorboard_run" "pike" {
  tensorboard_run_id = "pike-test-run"
  tensorboard         = "8176507224839946240"
  experiment          = "pike-dep-experiment"
  location            = "us-central1"
  display_name        = "pike-test-run"
}
