data "google_compute_instance_groups" "pike" {
  zone = "us-central1-b"
}

output "google_compute_instance_groups" {
  value = data.google_compute_instance_groups.pike
}
