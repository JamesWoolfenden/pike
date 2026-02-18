resource "google_compute_instance" "pike_source" {
  provider     = google-beta
  name         = "pike-source-instance"
  machine_type = "e2-micro"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }

  service_account {
    scopes = []
  }
}

resource "google_compute_machine_image" "pike" {
  provider        = google-beta
  name            = "pike-machine-image"
  source_instance = google_compute_instance.pike_source.self_link
}

resource "google_compute_instance_from_machine_image" "pike" {
  provider = google-beta
  name     = "pike-instance-from-image"
  zone     = "us-central1-a"

  source_machine_image = google_compute_machine_image.pike.self_link

  service_account {
    email  = ""
    scopes = []
  }
}
