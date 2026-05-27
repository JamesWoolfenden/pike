resource "google_pubsub_lite_reservation" "pike" {
  name                = "pike-reserve"
  throughput_capacity = 2
}
