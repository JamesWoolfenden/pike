resource "google_compute_instance_group_named_port" "pike" {
  zone  = "us-central1-a"
  group = google_compute_instance_group.pike.name
  name  = "http"
  port  = 8080
}
