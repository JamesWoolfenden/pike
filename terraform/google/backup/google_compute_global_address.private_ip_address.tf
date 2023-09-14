resource "google_compute_global_address" "private_ip_address" {
  name          = "pike"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = "https://www.googleapis.com/compute/v1/projects/pike-361314/global/networks/default"
}
