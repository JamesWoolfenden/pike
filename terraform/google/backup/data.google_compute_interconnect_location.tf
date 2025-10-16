data "google_compute_interconnect_location" "pike" {
  provider = google-beta
  name     = "pike"
}

output "google_compute_interconnect_location" {
  value = data.google_compute_interconnect_location.pike
}
