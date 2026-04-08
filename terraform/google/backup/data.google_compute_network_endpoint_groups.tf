data "google_compute_network_endpoint_groups" "pike" {
  zone = "us-central1-a"
}

output "google_compute_network_endpoint_groups" {
  value = data.google_compute_network_endpoint_groups.pike
}
