data "google_network_security_address_groups" "pike" {
  location = "us-central1"
}

output "google_network_security_address_groups" {
  value = data.google_network_security_address_groups.pike
}
