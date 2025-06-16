data "google_compute_images" "pike" {
}

output "google_compute_images" {
  value = data.google_compute_images.pike
}
