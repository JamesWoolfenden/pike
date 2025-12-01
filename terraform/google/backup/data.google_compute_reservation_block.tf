data "google_compute_reservation_block" "pike" {
  name        = "pike"
  reservation = "pike"
  zone        = "us-central1-a"
}

output "google_compute_reservation_block" {
  value = data.google_compute_reservation_block.pike
}
