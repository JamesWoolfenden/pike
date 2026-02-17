resource "google_compute_network" "pike_mirror_network" {
  name                    = "pike-mirror-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "pike_mirror_subnet" {
  name          = "pike-mirror-subnet"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.pike_mirror_network.id
}

resource "google_compute_instance" "pike_mirror_source" {
  name         = "pike-mirror-source"
  machine_type = "e2-micro"
  zone         = "us-central1-a"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    subnetwork = google_compute_subnetwork.pike_mirror_subnet.id
  }
}

resource "google_compute_health_check" "pike_mirror_health" {
  name = "pike-mirror-health"

  tcp_health_check {
    port = 80
  }
}

resource "google_compute_region_backend_service" "pike_mirror_backend" {
  name          = "pike-mirror-backend"
  region        = "us-central1"
  health_checks = [google_compute_health_check.pike_mirror_health.id]
}

resource "google_compute_forwarding_rule" "pike_mirror_ilb" {
  name                   = "pike-mirror-ilb"
  region                 = "us-central1"
  load_balancing_scheme  = "INTERNAL"
  backend_service        = google_compute_region_backend_service.pike_mirror_backend.id
  subnetwork             = google_compute_subnetwork.pike_mirror_subnet.id
  ip_protocol            = "TCP"
  ports                  = ["80"]
  is_mirroring_collector = true
}

resource "google_compute_packet_mirroring" "pike" {
  name   = "pike-packet-mirror"
  region = "us-central1"

  network {
    url = google_compute_network.pike_mirror_network.id
  }

  collector_ilb {
    url = google_compute_forwarding_rule.pike_mirror_ilb.id
  }

  mirrored_resources {
    instances {
      url = google_compute_instance.pike_mirror_source.id
    }
  }
}
