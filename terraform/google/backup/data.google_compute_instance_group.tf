data "google_compute_instance_group" "pike" {
  name = "pike"
  zone = "us-central1-b"
}

output "group" {
  value = data.google_compute_instance_group.pike
}
