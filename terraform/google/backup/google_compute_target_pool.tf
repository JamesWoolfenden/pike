
resource "google_compute_target_pool" "pike" {
  name   = "instance-pool"
  region = "us-central1"
  instances = [
    "us-central1-a/${google_compute_instance.pike.name}"
  ]

  health_checks = [
    google_compute_http_health_check.default.name,
  ]
}


resource "google_compute_instance" "pike" {
  machine_type = "n2-standard-2"
  name         = "pike"
  network_interface {
    network = google_compute_network.temp.name
  }

  allow_stopping_for_update = true
  zone                      = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
      labels = {
        my_label = "value"
      }
    }
  }
}

resource "google_compute_http_health_check" "default" {
  name               = "default"
  request_path       = "/"
  check_interval_sec = 1
  timeout_sec        = 1
}
