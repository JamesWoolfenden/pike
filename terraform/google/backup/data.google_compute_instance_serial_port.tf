data "google_compute_instance_serial_port" "pike" {
  instance = "my-instance"
  zone     = "us-central1-a"
  port     = 1
}
