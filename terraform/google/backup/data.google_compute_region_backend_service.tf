data "google_compute_region_backend_service" "pike" {
  name = "pike"
}

output "google_compute_region_backend_service" {
  value = data.google_compute_region_backend_service.pike
}
