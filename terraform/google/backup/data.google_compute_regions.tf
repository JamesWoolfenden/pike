data "google_compute_regions" "pike" {}

output "google_compute_regions" {
  value = data.google_compute_regions.pike
}
