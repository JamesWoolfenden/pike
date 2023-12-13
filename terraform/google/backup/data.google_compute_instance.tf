data "google_compute_instance" "pike" {
  zone = "us-central1-b"
  name = "pike"
}

output "instance" {
  value = data.google_compute_instance.pike
}
