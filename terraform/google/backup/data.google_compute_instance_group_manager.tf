data "google_compute_instance_group_manager" "pike" {
  name = "my-igm"
  zone = "us-central1-a"
}
