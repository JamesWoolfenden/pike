data "google_network_services_gateway" "pike" {
  name     = "pike-test-gateway"
  location = "europe-west2"
}

output "google_network_services_gateway" {
  value = data.google_network_services_gateway.pike
}