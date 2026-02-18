resource "google_compute_disk" "pike_primary" {
  name = "pike-primary-disk"
  type = "pd-balanced"
  zone = "us-central1-a"
  size = 10
}

resource "google_compute_disk" "pike_secondary" {
  name = "pike-secondary-disk"
  type = "pd-balanced"
  zone = "us-west1-a"
  size = 10
  async_primary_disk {
    disk = google_compute_disk.pike_primary.id
  }
}

resource "google_compute_disk_async_replication" "pike" {
  primary_disk = google_compute_disk.pike_primary.id
  secondary_disk {
    disk = google_compute_disk.pike_secondary.id
  }
}
