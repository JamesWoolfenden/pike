resource "google_lustre_instance" "pike" {
  instance_id                 = "my-instance"
  location                    = "us-central1-a"
  description                 = "test lustre instance"
  filesystem                  = "testfs"
  per_unit_storage_throughput = "125"
  capacity_gib                = 72000
  network                     = data.google_compute_network.lustre-network.id
  labels = {
    test = "value"
  }
  timeouts {
    create = "120m"
  }
}

data "google_compute_network" "lustre-network" {
  name = "default"
}
