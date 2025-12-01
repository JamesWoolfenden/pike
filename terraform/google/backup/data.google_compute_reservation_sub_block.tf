data "google_compute_reservation_sub_block" "pike" {
  name              = "pike"
  reservation_block = "pike"
  reservation       = "pike"
  zone              = "us-central1-a"
}

output "google_compute_reservation_sub_block" {
  value = data.google_compute_reservation_sub_block.pike
}
