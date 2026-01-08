resource "google_data_fusion_instance" "basic_instance" {
  name   = "my-instance"
  region = "us-central1"
  type   = "BASIC"
}
