resource "google_compute_address" "static" {
  name = "bastion"
  #  network_tier = "PREMIUM"
  //  address_type = "INTERNAL"
  description = "bastion IP"
  //network = "pike"
  //  subnetwork = "pike"
  # region       = var.region
}
