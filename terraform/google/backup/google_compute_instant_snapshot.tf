resource "google_compute_disk" "pike_snapshot_disk" {
  name = "pike-snapshot-disk"
  type = "pd-standard"
  zone = "us-central1-a"
  size = 10
}

resource "google_compute_instant_snapshot" "pike" {
  name        = "pike-instant-snapshot"
  zone        = "us-central1-a"
  source_disk = google_compute_disk.pike_snapshot_disk.id
}
