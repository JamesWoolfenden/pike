data "google_compute_subnetworks" "pike" {
}

output "google_compute_subnetworks" {
  value = data.google_compute_subnetworks.pike
}
