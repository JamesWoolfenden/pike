resource "google_compute_disk" "pike_attached" {
  name = "pike-attached-disk"
  type = "pd-standard"
  zone = "us-central1-a"
  size = 10
}

resource "google_compute_instance" "pike_attached" {
  name         = "pike-attached-instance"
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
}

resource "google_compute_attached_disk" "pike" {
  disk     = google_compute_disk.pike_attached.id
  instance = google_compute_instance.pike_attached.id
}
