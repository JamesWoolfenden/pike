resource "google_compute_global_address" "pike_private_ip_alloc" {
  name          = "pike-private-ip-alloc"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.pike.id
}

resource "google_service_networking_connection" "pike" {
  network                 = google_compute_network.pike.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.pike_private_ip_alloc.name]
}
