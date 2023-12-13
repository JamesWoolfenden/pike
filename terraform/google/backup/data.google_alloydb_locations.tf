data "google_alloydb_locations" "pike" {}

output "locations" {
  value = data.google_alloydb_locations.pike
}
