data "google_compute_instance_guest_attributes" "pike" {
  name = "pike"
  zone = "us-central1-a"
}

output "google_compute_instance_guest_attributes" {
  value = data.google_compute_instance_guest_attributes.pike
}
