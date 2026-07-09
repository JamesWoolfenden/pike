resource "google_compute_region_disk" "pike" {
  name          = "pike"
  region        = "us-central1"
  replica_zones = ["us-central1-a", "us-central1-b"]
  size          = 10
  type          = "pd-ssd"
}
