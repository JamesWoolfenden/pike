resource "google_compute_reservation" "pike" {
  name = "pike-reservation"
  zone = "us-central1-a"

  specific_reservation {
    count = 1
    instance_properties {
      machine_type = "n1-standard-1"
    }
  }
}
