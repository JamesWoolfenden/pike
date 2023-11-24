data "google_compute_network_peering" "pike" {
  name    = "pike"
  network = "projects/pike/global/networks/abc"
}
