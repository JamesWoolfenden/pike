data "google_compute_service_attachment" "pike" {
  name   = "pike"
  region = "us-central1"
}

output "google_compute_service_attachment" {
  value = data.google_compute_service_attachment.pike
}
