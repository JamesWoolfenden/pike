# Resource Policy Attachment - attaches a resource policy to an instance
resource "google_compute_resource_policy" "pike_policy" {
  name    = "pike-resource-policy"
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

resource "google_compute_instance" "pike_instance" {
  name         = "pike-instance-for-policy"
  machine_type = "e2-micro"
  zone         = "us-central1-a"
  project      = "pike-477416"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_resource_policy_attachment" "pike" {
  name     = google_compute_resource_policy.pike_policy.name
  instance = google_compute_instance.pike_instance.name
  zone     = "us-central1-a"
  project  = "pike-477416"
}
