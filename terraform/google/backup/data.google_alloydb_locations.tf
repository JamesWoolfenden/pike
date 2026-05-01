data "google_alloydb_locations" "pike" {}

output "google_alloydb_locations" {
  value = data.google_alloydb_locations.pike
}
