# Resource policy for disk scheduling
resource "google_compute_resource_policy" "pike_policy" {
  name   = "pike-disk-policy"
  region = "us-central1"

  snapshot_schedule_policy {
    schedule {
      daily_schedule {
        days_in_cycle = 1
        start_time    = "04:00"
      }
    }
  }
}

# Disk to attach policy to
resource "google_compute_disk" "pike_policy_disk" {
  name = "pike-policy-disk"
  type = "pd-standard"
  zone = "us-central1-a"
  size = 10
}

# Attachment resource - attaches policy to disk
resource "google_compute_disk_resource_policy_attachment" "pike" {
  name = google_compute_resource_policy.pike_policy.name
  disk = google_compute_disk.pike_policy_disk.name
  zone = "us-central1-a"
}
