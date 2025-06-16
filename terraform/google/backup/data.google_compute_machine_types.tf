data "google_compute_machine_types" "pike" {
  zone = "us-central1-a"
}

output "google_compute_machine_types" {
  value = data.google_compute_machine_types.pike
}
