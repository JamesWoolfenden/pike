resource "google_compute_region_instance_template" "pike" {
  name         = "pike-region-instance-template"
  region       = "europe-west2"
  project      = "pike-477416"
  machine_type = "e2-medium"
  description  = "Pike test template"

  disk {
    source_image = "debian-cloud/debian-11"
    auto_delete  = true
    boot         = true
  }

  network_interface {
    network = "default"
  }
}
