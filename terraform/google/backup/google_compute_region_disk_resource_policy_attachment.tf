# Region Disk Resource Policy Attachment - attaches a resource policy to a regional disk
resource "google_compute_resource_policy" "pike_policy" {
  name    = "pike-disk-policy"
  region  = "us-central1"
  project = "pike-477416"

  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time    = "04:00"
      }
    }
  }
}

resource "google_compute_region_disk" "pike_disk" {
  name    = "pike-region-disk"
  region  = "us-central1"
  project = "pike-477416"
  type    = "pd-standard"
  size    = 200

  replica_zones = ["us-central1-a", "us-central1-b"]
}

resource "google_compute_region_disk_resource_policy_attachment" "pike" {
  name    = google_compute_resource_policy.pike_policy.name
  disk    = google_compute_region_disk.pike_disk.name
  region  = "us-central1"
  project = "pike-477416"
}
