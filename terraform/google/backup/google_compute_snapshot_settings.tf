resource "google_compute_snapshot_settings" "pike" {
  storage_location {
    policy = "SPECIFIC_LOCATIONS"
    locations {
      name     = "us-central1"
      location = "us-central1"
    }
  }
}
