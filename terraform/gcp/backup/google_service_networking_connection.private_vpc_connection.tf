resource "google_service_networking_connection" "private_vpc_connection" {
  network                 = "https://www.googleapis.com/compute/v1/projects/pike-361314/global/networks/default"
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = ["pike"]
}
