data "google_compute_interconnect_locations" "pike" {
  provider = google-beta
}

output "google_compute_interconnect_locations" {
  value = data.google_compute_interconnect_locations.pike
}
