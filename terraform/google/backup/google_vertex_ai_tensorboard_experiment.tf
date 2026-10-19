resource "google_vertex_ai_tensorboard_experiment" "pike" {
  tensorboard_experiment_id = "pike-test-experiment"
  tensorboard               = "projects/430943803039/locations/us-central1/tensorboards/3186518837713436672"
  location                  = "us-central1"
}
