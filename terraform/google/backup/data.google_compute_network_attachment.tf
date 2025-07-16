data "google_compute_network_attachment" "pike" {
}

output "google_compute_network_attachment" {
  value = data.google_compute_network_attachment.pike
}
