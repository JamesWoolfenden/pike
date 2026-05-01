data "google_compute_subnetwork" "pike" {
  name = "pike"
}

output "google_compute_subnetwork" {
  value = data.google_compute_subnetwork.pike
}
