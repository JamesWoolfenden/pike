data "google_compute_zones" "all" {}

output "google_compute_zones" {
  value = data.google_compute_zones.all
}
