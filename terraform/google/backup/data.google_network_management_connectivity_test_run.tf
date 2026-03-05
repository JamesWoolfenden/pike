data "google_network_management_connectivity_test_run" "pike" {
  name = "pike"
}

output "google_network_management_connectivity_test_run" {
  value = data.google_network_management_connectivity_test_run.pike
}
