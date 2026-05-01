data "google_compute_image" "image" {
  family  = "debian-11"
  project = "debian-cloud"
}

output "google_compute_image_image" {
  value = data.google_compute_image.image
}
