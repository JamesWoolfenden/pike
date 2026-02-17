resource "google_compute_instance_group" "pike" {
  name        = "terraform-test"
  description = "Terraform test instance group"
  zone        = "us-central1-a"
  network     = google_compute_network.temp.id
}


resource "google_compute_network" "temp" {
  name = "temp"

}
