resource "google_compute_instance" "pike_instance" {
  name         = "pike-target-test"
  machine_type = "e2-micro"
  zone         = "europe-west2-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_target_instance" "pike" {
  name     = "pike-target-instance"
  instance = google_compute_instance.pike_instance.id
  zone     = "europe-west2-a"
}
