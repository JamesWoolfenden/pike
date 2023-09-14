data "google_compute_image" "image" {
  family  = "debian-11"
  project = "debian-cloud"
}

output "image" {
  value = data.google_compute_image.image
}
