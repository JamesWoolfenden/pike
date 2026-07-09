resource "google_compute_region_instant_snapshot" "pike" {
  name        = "pike"
  region      = "us-central1"
  source_disk = google_compute_region_disk.pike.id
}
