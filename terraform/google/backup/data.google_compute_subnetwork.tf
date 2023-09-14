data "google_compute_subnetwork" "pike" {
  name = "pike"
}

output "subnetwork" {
  value = data.google_compute_subnetwork.pike
}
