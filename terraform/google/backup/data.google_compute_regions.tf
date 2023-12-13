data "google_compute_regions" "pike" {}

output "regions" {
  value = data.google_compute_regions.pike
}
