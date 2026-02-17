resource "google_compute_global_network_endpoint_group" "pike_group" {
  name                  = "pike-global-neg-test"
  network_endpoint_type = "INTERNET_FQDN_PORT"
  default_port          = 443
}

resource "google_compute_global_network_endpoint" "pike" {
  global_network_endpoint_group = google_compute_global_network_endpoint_group.pike_group.name
  fqdn                          = "example.com"
  port                          = 443
}
