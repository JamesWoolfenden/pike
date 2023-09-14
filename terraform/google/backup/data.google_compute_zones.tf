data "google_compute_zones" "all" {}

output "zones" {
  value = data.google_compute_zones.all
}
