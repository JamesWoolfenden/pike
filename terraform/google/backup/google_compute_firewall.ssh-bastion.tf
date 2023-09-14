resource "google_compute_firewall" "ssh-bastion" {
  name        = "pike"
  description = "firewall to bastion"
  network     = "pike"
  // project     = var.network_interface["subnetwork_project"]

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = ["10.0.0.0/16"]
  //target_tags   = [google_compute_instance.bastion.name]

  log_config {
    metadata = "INCLUDE_ALL_METADATA"
  }
}
