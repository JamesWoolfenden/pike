resource "google_compute_instance_template" "pike" {
  name         = "pike-instance-template"
  machine_type = "e2-micro"

  disk {
    source_image = "debian-cloud/debian-11"
    auto_delete  = true
    boot         = true
  }

  network_interface {
    network = "default"
  }
}

resource "google_compute_instance_from_template" "pike" {
  name = "pike-instance-from-template"
  zone = "us-central1-a"

  source_instance_template = google_compute_instance_template.pike.self_link

  labels = {
    environment = "test"
  }
}
