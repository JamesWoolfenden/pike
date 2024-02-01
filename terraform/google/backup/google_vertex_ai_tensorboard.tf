resource "google_vertex_ai_tensorboard" "pike" {
  display_name = "terraform"
  description  = "sample description"
  labels = {
    "key1" : "value1",
    "key2" : "value2"
  }
  region = "us-central1"
}
