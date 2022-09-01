resource "google_compute_instance" "default" {
  name         = "test"
  machine_type = "e2-micro"
  zone         = "europe-west2-a"



  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  // Local SSD disk
  #  scratch_disk {
  #    interface = "SCSI"
  #  }

  network_interface {
    network = "default"

    access_config {
      // Ephemeral public IP
    }
  }

  metadata = {
    foo = "bar"
  }

  metadata_startup_script = "echo hi > /test.txt"

  #  service_account {
  #    # Google recommends custom service accounts that have cloud-platform scope and permissions granted via IAM Roles.
  #    email  = google_service_account.default.email
  #    scopes = ["cloud-platform"]
  #  }

  tags = ["foo", "bar"]
}
