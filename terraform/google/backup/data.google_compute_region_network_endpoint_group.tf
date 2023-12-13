data "google_compute_region_network_endpoint_group" "pike" {
  name = "pike"
}

output "group" {
  value = data.google_compute_region_network_endpoint_group.pike
}
