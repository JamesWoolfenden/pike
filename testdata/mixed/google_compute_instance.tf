resource "google_compute_instance" "pike" {
  machine_type = ""
  name         = "pike"
  boot_disk {}
  network_interface {}
}
