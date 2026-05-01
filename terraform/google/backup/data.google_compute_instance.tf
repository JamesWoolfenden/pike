data "google_compute_instance" "pike" {
  zone = "us-central1-b"
  name = "pike"
}

output "google_compute_instance" {
  value = data.google_compute_instance.pike
}
