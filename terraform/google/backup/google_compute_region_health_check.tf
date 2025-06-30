resource "google_compute_region_health_check" "pike" {
  name = "tcp-region-health-check"

  timeout_sec        = 1
  check_interval_sec = 1

  ssl_health_check {
    port = "443"
  }
  region = "us-central1"
}
