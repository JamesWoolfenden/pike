data "google_network_management_connectivity_tests" "pike" {
}

output "google_network_management_connectivity_tests" {
  value = data.google_network_management_connectivity_tests.pike
}
